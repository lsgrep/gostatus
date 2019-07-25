package addon

import (
	"context"
	"fmt"
	"github.com/lsgrep/gostatus/log"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type ipv6Ext struct {
}

// single stack
var dialerV6 = &net.Dialer{DualStack: false}

var dialContextV6 = func(ctx context.Context, network, addr string) (net.Conn, error) {
	network = "tcp6"
	return dialerV6.DialContext(ctx, network, addr)
}

var httpClientV6 = &http.Client{
	Transport: &http.Transport{
		Proxy:       http.ProxyFromEnvironment,
		DialContext: dialContextV6,
	},
}

// Get preferred outbound ip of this machine
func (i *ipv6Ext) Update() *Block {
	resp, err := httpClientV6.Get("https://api6.ipify.org")
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

func NewIpv6ExtAddon() *Addon {
	i := &ipv6Ext{}
	aa := Addon{
		UpdateInterval: 5000 * time.Millisecond,
		Updater:        i,
	}
	return &aa
}
