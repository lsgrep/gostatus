package addon

import (
	"fmt"
	"io/ioutil"
	"time"
)

type ipExt struct {
}

func (ie *ipExt) Update() *Block {
	resp, err := httpCli.Get("https://api.ipify.org")
	if err != nil {
		logger.Error(err)
		return nil
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
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
