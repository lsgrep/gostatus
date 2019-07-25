package bar

import (
	"fmt"
	"time"

	"github.com/lsgrep/gostatus/config"

	"encoding/json"

	"bytes"

	"bufio"
	"os"

	"github.com/lsgrep/gostatus/addon"
	"github.com/lsgrep/gostatus/log"
)

// https://i3wm.org/docs/i3bar-protocol.html
var initMsg = `{ "version": 1, "stop_signal": 10, "cont_signal": 12, "click_events": true }`

type gostatus struct {
	addons []*addon.Addon
}

type Bar interface {
	Run(configPath string)
	Add(addon *addon.Addon)
}

func setupBar() {
	fmt.Print(initMsg)
	// let's start the endless array
	fmt.Print("[")

	// first array as empty
	fmt.Print("[]")
}

func (gs *gostatus) processInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		_, _, err := reader.ReadLine()
		if err != nil {
			log.Error(err)
			break
		}
	}
}

func (gs *gostatus) render() {
	buf := bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	for {
		buf.Reset()
		var output []addon.Block
		for _, a := range gs.addons {
			if a.LastData != nil {
				temp := *a.LastData
				output = append(output, temp)
			}
		}

		if len(output) == 0 {
			continue
		}

		err := encoder.Encode(output)
		if err != nil {
			log.Error(err)
			break
		}
		//necessary to start with a comma
		fmt.Print(",")
		fmt.Print(string(buf.Bytes()))
		time.Sleep(time.Second)
	}
}

func (gs *gostatus) Run(filePath string) {
	// 0. load config
	gs.loadConfig(filePath)

	// 1. setup i3bar
	setupBar()

	// 2. process events
	go gs.processInput()

	// 3. run addons
	for _, a := range gs.addons {
		go a.Run()
	}

	// 3. render addons
	gs.render()
}

func NewGoStatusBar() *gostatus {
	return &gostatus{}
}

func (gs *gostatus) Add(a *addon.Addon) {
	gs.addons = append(gs.addons, a)
}

func (gs *gostatus) loadConfig(filePath string) {
	addons, err := config.ReadConfig(filePath)
	if err != nil {
		gs.Add(addon.NewMessageAddon(err.Error()))
	}
	for _, a := range addons {
		gs.Add(a)
	}
}
