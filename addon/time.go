package addon

import "time"

type timer struct {
}

func (t *timer) Update() string {
	return time.Now().Format(time.RFC822)
}

func NewTimeAddon() *Addon {
	t := &timer{}
	aa := Addon{
		Icon:             "\uf017",
		UpdateIntervalMs: 1000,
		Updater:          t}
	return &aa
}
