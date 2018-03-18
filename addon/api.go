package addon

import (
	"sync"
	"time"
)

// for each data fetcher
type Updater interface {
	Update() string
}

type Addon struct {
	sync.Mutex
	UpdateIntervalMs int64
	Icon             string
	Updater          Updater
	LastData         *Block
}

func (a *Addon) Run() {
	a.LastData = &Block{}
	for range time.Tick(time.Duration(int64(time.Millisecond) * a.UpdateIntervalMs)) {
		a.Lock()
		data := a.Updater.Update()
		if data == "" {
			a.LastData.FullText = ""
		} else {
			a.LastData.FullText = a.Icon + " " + a.Updater.Update()
		}
		a.Unlock()
	}
}
