package addon

import (
	"fmt"
	"github.com/lsgrep/gostatus/log"
	"net"
	"time"
)

type ip struct {
	networkInterface string
}

// Get preferred outbound ip of this machine
func (i *ip) Update() *Block {
	var ip net.IP
	iface, err := net.InterfaceByName(i.networkInterface)
	if err != nil {
		log.Error(err)
		return nil
	}

	addrs, err := iface.Addrs()
	if err != nil {
		log.Error(err)
		return nil
	}
	// handle err
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPNet:
			if v.IP.To4() != nil {
				ip = v.IP
			}
		}
		// process IP address
	}
	fullTxt := fmt.Sprintf(" %s  %s", IconIP, ip.String())
	return &Block{FullText: fullTxt, Color: ColorLime}
}

func NewIpAddon(iface string) *Addon {
	i := &ip{networkInterface: iface}
	aa := Addon{
		UpdateInterval: 5000 * time.Millisecond,
		Updater:        i,
	}
	return &aa
}
