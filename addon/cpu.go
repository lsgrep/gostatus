package addon

import (
	"fmt"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func GetCPUInfo() string {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		panic(err)
	}

	cpu := stat.CPUStatAll
	cpuUsage := cpu.System + cpu.User + cpu.Nice + cpu.Guest + cpu.GuestNice + cpu.IOWait + cpu.IRQ + cpu.SoftIRQ + cpu.Steal
	usage := float64(cpuUsage*100) / float64(cpuUsage+cpu.Idle)
	return fmt.Sprintf("%.2f%%", usage)
}

func NewCPUAddon() *Addon {
	return &Addon{UpdateIntervalMs: 3000,
		UpdateFn: func(a *Addon) {
			a.LastData = &Block{FullText: GetCPUInfo()}
		}}
}
