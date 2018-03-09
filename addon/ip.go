package addon

import (
	"net"

	"github.com/monax/monax/log"
)

// Get preferred outbound ip of this machine
func GetOutboundIP(networkInterface string) net.IP {
	var ip net.IP
	iface, err := net.InterfaceByName(networkInterface)
	if err != nil {
		log.Panic(err)
	}

	addrs, err := iface.Addrs()
	if err != nil {
		log.Panic(err)
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
	return ip
}

func NewIpAddon(iface string) *Addon {
	aa := Addon{UpdateIntervalMs: 5000,
		UpdateFn: func(a *Addon) {
			a.LastData = &Block{FullText: GetOutboundIP(iface).String()}
		}}
	return &aa
}
