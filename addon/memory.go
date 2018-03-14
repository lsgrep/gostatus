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

func GetMemory() (int64, int64) {
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

	return int64(available), int64(total)
}

func NewMemoryAddon() *Addon {
	return &Addon{UpdateIntervalMs: 3000,
		UpdateFn: func(a *Addon) {
			avail, total := GetMemory()
			a.LastData = &Block{
				FullText: "\uf2db " + fmt.Sprintf("%.2fGB / %.2fGB",
					float64(avail)/1024/1024, float64(total)/1024/1024)}
		}}
}
