package addon

import (
	"fmt"
	"io/ioutil"

	"strings"

	"strconv"
)

/*
MemTotal:       16338116 kB
MemFree:          215604 kB
MemAvailable:    9196056 kB
*/

type memory struct {
}

func (m *memory) Update() string {
	var err error
	buf, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(buf), "\n")
	memTotal := strings.Fields(lines[0])[1]
	memAvailable := strings.Fields(lines[2])[1]

	total, err := strconv.Atoi(memTotal)
	if err != nil {
		panic(err)
	}

	available, err := strconv.Atoi(memAvailable)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%.2fGB / %.2fGB",
		float64(available)/1024/1024, float64(total)/1024/1024)
}

func NewMemoryAddon() *Addon {
	m := &memory{}
	return &Addon{
		UpdateIntervalMs: 3000,
		Icon:             "\uf2db",
		Updater:          m}
}
