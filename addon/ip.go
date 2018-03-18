package addon

import (
	"net"
)

type ip struct {
	networkInterface string
}

// Get preferred outbound ip of this machine
func (i *ip) Update() string {
	var ip net.IP
	iface, err := net.InterfaceByName(i.networkInterface)
	if err != nil {
		panic(err)
	}

	addrs, err := iface.Addrs()
	if err != nil {
		panic(err)
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
	return ip.String()
}

func NewIpAddon(iface string) *Addon {
	i := &ip{networkInterface: iface}
	aa := Addon{
		UpdateIntervalMs: 5000,
		Icon:             "\uf0e8",
		Updater:          i,
	}
	return &aa
}
