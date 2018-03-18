package main

import (
	"github.com/lsgrep/gostatus/addon"
	"github.com/lsgrep/gostatus/bar"
)

func main() {

	statusBar := bar.NewGoStatusBar()
	statusBar.Add(addon.NewVolumeAddon())
	statusBar.Add(addon.NewNetworkAddon("enp5s0"))
	statusBar.Add(addon.NewGithubNotificationsAddon("lsgrep"))
	statusBar.Add(addon.NewCPUAddon())

	// pass network interface name
	statusBar.Add(addon.NewIpAddon("enp5s0"))
	statusBar.Add(addon.NewMemoryAddon())
	statusBar.Add(addon.NewDiskAddon("/data"))
	statusBar.Add(addon.NewTimeAddon())
	statusBar.Run()
}
