package addon

import (
	"sync"
	"time"
)

type Updater interface {
	Update()
	LastData() *Block
}

type Addon struct {
	sync.Mutex
	UpdateIntervalMs int64
	UpdateFn         func(addon *Addon)
	LastData         *Block
}

func (a *Addon) Update() {
	for range time.Tick(time.Duration(int64(time.Millisecond) * a.UpdateIntervalMs)) {
		a.Lock()
		a.UpdateFn(a)
		a.Unlock()
	}
}
