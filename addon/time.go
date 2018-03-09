package addon

import "time"

func NewTimeAddon() *Addon {
	aa := Addon{UpdateIntervalMs: 1000,
		UpdateFn: func(a *Addon) {
			a.LastData = &Block{FullText: time.Now().Format(time.RFC822)}
		}}
	return &aa
}
