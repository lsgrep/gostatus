package addon

import (
	"net"
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

/*
"color": "#00ff00",
    "background": "#0000ff"
*/
func NewIpAddon(iface string) *Addon {
	aa := Addon{UpdateIntervalMs: 5000,
		UpdateFn: func(a *Addon) {
			ip := GetOutboundIP(iface).String()
			a.LastData = &Block{FullText: "\uf0e8" + ip, Color: "#00ff00"}
		}}
	return &aa
}
