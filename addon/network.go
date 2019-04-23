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

func GetNetwork(iface string) (int64, int64, error) {
	var err error
	buf, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		logger.Error(err)
		return 0, 0, err
	}

	data := strings.Split(string(buf), "\n")
	var ifaceData string
	for _, l := range data {
		if strings.HasPrefix(l, iface) {
			ifaceData = l
		}
	}

	fields := strings.Fields(ifaceData)
	downCount, _ := strconv.Atoi(fields[1])
	uploadCount, _ := strconv.Atoi(fields[9])
	return int64(downCount), int64(uploadCount), nil
}

func (ns *networkStatus) Update() *Block {
	ts := time.Now().Unix()
	downCnt, upCnt, err := GetNetwork(ns.NetworkInterface)
	if err != nil {
		logger.Error(err)
		return nil
	}
	if ns.DownPacketCnt == 0 || ns.UpPacketCnt == 0 {
		ns.checkedAt = ts
		ns.DownPacketCnt = downCnt
		ns.UpPacketCnt = upCnt
		return nil
	}
	downSpeed := float64(downCnt-ns.DownPacketCnt) / float64(ts-ns.checkedAt) / 1024
	upSpeed := float64(upCnt-ns.UpPacketCnt) / float64(ts-ns.checkedAt) / 1024

	txt := fmt.Sprintf("%.2f KB/s %.2f KB/s", downSpeed, upSpeed)
	fullTxt := fmt.Sprintf(" %s  %s", IconNetwork, txt)

	return &Block{FullText: fullTxt}
}

func NewNetworkAddon(iface string) *Addon {
	n := &networkStatus{NetworkInterface: iface}
	return &Addon{
		UpdateInterval: 3000 * time.Millisecond,
		Updater:        n}
}
