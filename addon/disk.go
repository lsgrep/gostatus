package addon

import (
	"fmt"
	"syscall"
)

type DiskStatus struct {
	All  uint64
	Free uint64
	Used uint64
}

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func NewDiskAddon(path string) *Addon {
	return &Addon{UpdateIntervalMs: 5000,
		UpdateFn: func(a *Addon) {
			disk := DiskUsage(path)
			a.LastData = &Block{FullText: fmt.Sprintf("\uf1c0  %s %.2fGB / %.2fGB", path,
				float64(disk.Free)/float64(GB), float64(disk.All)/float64(GB))}
		}}
}
