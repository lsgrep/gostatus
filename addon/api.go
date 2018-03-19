package addon

import (
	"sync"
	"time"
)

// for each data fetcher
type Updater interface {
	Update() *Block
}

type Addon struct {
	// guard lastData
	sync.Mutex
	LastData *Block

	UpdateIntervalMs int64
	Icon             string
	Updater          Updater
}

func (a *Addon) Run() {
	for range time.Tick(time.Duration(int64(time.Millisecond) * a.UpdateIntervalMs)) {
		a.Lock()
		a.LastData = a.Updater.Update()
		a.Unlock()
	}
}
