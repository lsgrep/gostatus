package bar

import (
	"fmt"
	"time"

	"encoding/json"

	"bytes"

	"bufio"
	"os"

	"github.com/apex/log"
	"github.com/dexDev/utils/ulog"
	"github.com/lsgrep/gostatus/addon"
)

// https://i3wm.org/docs/i3bar-protocol.html
var initMsg = `{ "version": 1, "stop_signal": 10, "cont_signal": 12, "click_events": true }`

type gostatus struct {
	log    ulog.Logger
	addons []*addon.Addon
}

type Bar interface {
	Run()
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
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		gs.log.Info(isPrefix, string(line))
	}
}

func (gs *gostatus) render() {
	buf := bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	//encoder.SetIndent("", " ")
	for {
		buf.Reset()
		var output []addon.Block
		for _, a := range gs.addons {
			a.Lock()
			if a.LastData != nil {
				temp := *a.LastData
				output = append(output, temp)
			}
			a.Unlock()
		}

		if len(output) == 0 {
			continue
		}

		err := encoder.Encode(output)
		if err != nil {
			panic(err)
		}
		//necessary to start with a comma
		fmt.Print(",")
		fmt.Print(string(buf.Bytes()))
		gs.log.Info(string(buf.Bytes()))
		time.Sleep(time.Second)
	}
}

func (gs *gostatus) Run() {
	// 1. setup i3bar
	setupBar()

	// 2. process events
	go gs.processInput()

	// 3. run addons
	for _, a := range gs.addons {
		go a.Update()
	}

	log.Info("rendering...")
	// 3. render addons
	gs.render()
}

func NewGoStatusBar() *gostatus {
	return &gostatus{log: ulog.NewLogger()}
}

func (gs *gostatus) Add(a *addon.Addon) {
	gs.addons = append(gs.addons, a)
}
