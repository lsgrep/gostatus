package addon

import (
	"fmt"
	"syscall"
	"time"
)

type diskStatus struct {
	Path string
	All  uint64
	Free uint64
	Used uint64
}

func (ds *diskStatus) Update() *Block {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(ds.Path, &fs)
	if err != nil {
		return nil
	}
	ds.All = fs.Blocks * uint64(fs.Bsize)
	ds.Free = fs.Bfree * uint64(fs.Bsize)
	ds.Used = ds.All - ds.Free
	text := fmt.Sprintf("%s %.2fGB / %.2fGB", ds.Path,
		float64(ds.Free)/float64(GB), float64(ds.All)/float64(GB))
	fullTxt := fmt.Sprintf(" %s  %s", IconDisk, text)
	return &Block{FullText: fullTxt}
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func NewDiskAddon(path string) *Addon {
	ds := &diskStatus{Path: path}
	return &Addon{
		UpdateInterval: 5000 * time.Millisecond,
		Updater:        ds}
}
