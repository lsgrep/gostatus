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

func (m *memory) Update() *Block {
	var err error
	var memAvail, memTotal int64
	r, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil
	}
	defer r.Close()
	_, err = fmt.Fscanf(
		r,
		"MemTotal: %d kB\nMemFree: %d kB\nMemAvailable: %d ",
		&memTotal, &memAvail, &memAvail)

	txt := fmt.Sprintf("%.2fGB / %.2fGB",
		float64(memAvail)/1024/1024, float64(memTotal)/1024/1024)
	fullTxt := fmt.Sprintf(" %s  %s", IconMemory, txt)
	return &Block{FullText: fullTxt}
}

func NewMemoryAddon() *Addon {
	m := &memory{}
	return &Addon{
		UpdateIntervalMs: 3000,
		Updater:          m}
}
