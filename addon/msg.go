package addon

import (
	"time"
)

type message string

func (t message) Update() *Block {
	return &Block{FullText: string(t)}
}

func NewMessageAddon(m string) *Addon {
	aa := Addon{
		UpdateInterval: 1000 * time.Millisecond,
		Updater:        message(m)}
	return &aa
}
