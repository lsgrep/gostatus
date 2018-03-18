package addon

import (
	"fmt"

	"os"
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
	var memAvail, memTotal int64
	r, err := os.Open("/proc/meminfo")
	if err != nil {
		return ""
	}
	defer r.Close()
	_, err = fmt.Fscanf(
		r,
		"MemTotal: %d kB\nMemFree: %d kB\nMemAvailable: %d ",
		&memTotal, &memAvail, &memAvail)

	return fmt.Sprintf("%.2fGB / %.2fGB",
		float64(memAvail)/1024/1024, float64(memTotal)/1024/1024)
}

func NewMemoryAddon() *Addon {
	m := &memory{}
	return &Addon{
		UpdateIntervalMs: 3000,
		Icon:             "\uf2db",
		Updater:          m}
}
