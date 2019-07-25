package addon

import (
	"fmt"
	"github.com/lsgrep/gostatus/log"

	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type cpu struct {
}

func (c *cpu) Update() *Block {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Error(err)
		return nil
	}

	cpu := stat.CPUStatAll
	cpuUsage := cpu.System + cpu.User + cpu.Nice + cpu.Guest + cpu.GuestNice + cpu.IOWait + cpu.IRQ + cpu.SoftIRQ + cpu.Steal
	usage := float64(cpuUsage*100) / float64(cpuUsage+cpu.Idle)

	fullTxt := fmt.Sprintf(" %s  %s", IconCPU, fmt.Sprintf("%.2f%%", usage))
	b := &Block{FullText: fullTxt}
	return b
}

func NewCPUAddon() *Addon {
	c := &cpu{}
	return &Addon{
		UpdateInterval: 3000 * time.Millisecond,
		Updater:        c}
}
