package addon

import (
	"fmt"
	"syscall"
)

type DiskStatus struct {
	Path string
	All  uint64
	Free uint64
	Used uint64
}

func (ds *DiskStatus) Update() string {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(ds.Path, &fs)
	if err != nil {
		return ""
	}
	ds.All = fs.Blocks * uint64(fs.Bsize)
	ds.Free = fs.Bfree * uint64(fs.Bsize)
	ds.Used = ds.All - ds.Free
	return fmt.Sprintf("\uf1c0  %s %.2fGB / %.2fGB", ds.Path,
		float64(ds.Free)/float64(GB), float64(ds.All)/float64(GB))
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func NewDiskAddon(path string) *Addon {
	ds := &DiskStatus{Path: path}
	return &Addon{
		UpdateIntervalMs: 5000,
		Updater:          ds}
}
