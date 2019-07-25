package addon

import (
	"time"
)

// for each data fetcher
type Updater interface {
	Update() *Block
}

type Addon struct {
	// guard lastData
	LastData *Block

	UpdateInterval time.Duration
	Icon           string
	Updater        Updater
}

func (a *Addon) Run() {
	for range time.Tick(a.UpdateInterval) {
		// generating data should not be locked
		newData := a.Updater.Update()
		a.LastData = newData
	}
}
