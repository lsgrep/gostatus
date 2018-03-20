package addon

import (
	"fmt"

	"os"
	"time"
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
	return &Block{FullText: fullTxt, Color: ColorYellow}
}

func NewMemoryAddon() *Addon {
	m := &memory{}
	return &Addon{
		UpdateInterval: 3000 * time.Millisecond,
		Updater:        m}
}
