package addon

import (
	"fmt"
	"time"
)

type timer struct {
}

func (t *timer) Update() *Block {
	fullTxt := fmt.Sprintf(" %s  %s", IconTime, time.Now().Format(time.RFC822))
	return &Block{FullText: fullTxt}
}

func NewTimeAddon() *Addon {
	t := &timer{}
	aa := Addon{
		UpdateInterval: 1000 * time.Millisecond,
		Updater:        t}
	return &aa
}
