package addon

import (
	"io/ioutil"

	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
MemTotal:       16338116 kB
MemFree:          215604 kB
MemAvailable:    9196056 kB
*/

type networkStatus struct {
	NetworkInterface string
	DownPacketCnt    int64
	UpPacketCnt      int64
	checkedAt        int64
}

func GetNetwork(iface string) (int64, int64) {
	var err error
	buf, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(buf), "\n")
	var ifaceData string
	for _, l := range data {
		if strings.Contains(l, iface) {
			ifaceData = l
		}
	}

	fields := strings.Fields(ifaceData)
	downCount, _ := strconv.Atoi(fields[1])
	uploadCount, _ := strconv.Atoi(fields[9])
	return int64(downCount), int64(uploadCount)
}

func (ns *networkStatus) Update() string {
	ts := time.Now().Unix()
	downCnt, upCnt := GetNetwork(ns.NetworkInterface)
	if ns.DownPacketCnt == 0 || ns.UpPacketCnt == 0 {
		ns.checkedAt = ts
		ns.DownPacketCnt = downCnt
		ns.UpPacketCnt = upCnt
		return ""
	}
	downSpeed := float64(downCnt-ns.DownPacketCnt) / float64(ts-ns.checkedAt) / 1024
	upSpeed := float64(upCnt-ns.UpPacketCnt) / float64(ts-ns.checkedAt) / 1024

	return fmt.Sprintf("%.2f KB/s %.2f KB/s", downSpeed, upSpeed)
}

func NewNetworkAddon(iface string) *Addon {
	n := &networkStatus{NetworkInterface: iface}
	return &Addon{
		UpdateIntervalMs: 3000,
		Icon:             "\uf0c1",
		Updater:          n}
}
