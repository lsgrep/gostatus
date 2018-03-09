package addon

import (
	"os/exec"

	"bufio"
	"bytes"
	"fmt"

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
	cmd := exec.Command("sh", "-c", "cat /proc/meminfo")
	buf := &bytes.Buffer{}
	cmd.Stdout = bufio.NewWriter(buf)
	err = cmd.Run()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println()
	lines := strings.Split(buf.String(), "\n")
	memTotal := strings.Fields(lines[0])[1]
	memAvailable := strings.Fields(lines[2])[1]

	total, err := strconv.Atoi(memTotal)
	if err != nil {
		log.Panic(err)
	}

	available, err := strconv.Atoi(memAvailable)
	if err != nil {
		log.Panic(err)
	}

	return int64(available), int64(total)
}

func NewMemoryAddon() *Addon {
	return &Addon{UpdateIntervalMs: 3000,
		UpdateFn: func(a *Addon) {
			avail, total := GetMemory()
			a.LastData = &Block{FullText: fmt.Sprintf("%.2fGB / %.2fGB",
				float64(avail)/1024/1024, float64(total)/1024/1024)}
		}}
}
