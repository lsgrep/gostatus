package addon

import (
	"sync"
	"time"
	"github.com/lsgrep/gostatus/utils"
)
var logger = utils.NewLogger()

// for each data fetcher
type Updater interface {
	Update() *Block
}

type Addon struct {
	// guard lastData
	sync.Mutex
	LastData *Block

	UpdateInterval time.Duration
	Icon           string
	Updater        Updater
}

func (a *Addon) Run() {
	for range time.Tick(a.UpdateInterval) {
		// generating data should not be locked
		newData := a.Updater.Update()

		a.Lock()
		a.LastData = newData
		a.Unlock()
	}
}
