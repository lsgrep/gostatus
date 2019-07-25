package addon

import (
	"fmt"
	"github.com/lsgrep/gostatus/log"
	"io/ioutil"
	"time"
)

type ipExt struct {
}

func (ie *ipExt) Update() *Block {
	resp, err := httpCli.Get("https://api.ipify.org")
	if err != nil {
		log.Error(err)
		return nil
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil
	}

	fullTxt := fmt.Sprintf(" %s %s", IconIP, string(bs))
	return &Block{FullText: fullTxt, Color: ColorLime}
}

func NewIpExtAddon() *Addon {
	ie := &ipExt{}
	ao := Addon{
		UpdateInterval: 10000 * time.Millisecond,
		Updater:        ie,
	}
	return &ao
}
