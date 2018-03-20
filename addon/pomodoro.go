package addon

import (
	"fmt"
	"time"
)

const (
	STOPPED = iota
	RUNNING_TASK
	RUNNING_BREAK
)

const (
	TaskDuration  = 25 * time.Minute
	BreakDuration = 5 * time.Minute
)

type pomodoro struct {
	State          int
	TaskRemaining  time.Duration
	BreakRemaining time.Duration
	UpdateInterval time.Duration
}

func (p *pomodoro) Update() *Block {
	var txt string
	if p.State == STOPPED {
		txt = "stopped"
	}

	if p.State == RUNNING_TASK {
		if p.TaskRemaining <= 0 {
			p.State = RUNNING_BREAK
			p.BreakRemaining = BreakDuration

			Notify("Take a Break!")
		} else {
			p.TaskRemaining -= p.UpdateInterval

		}
	}

	if p.State == RUNNING_BREAK {
		if p.BreakRemaining <= 0 {
			p.State = RUNNING_TASK
			p.TaskRemaining = TaskDuration

			Notify("Back to Work!")
		} else {
			p.BreakRemaining -= p.UpdateInterval
		}
	}

	if p.State == RUNNING_TASK {
		txt = IconWork + "  " + fmtDuration(p.TaskRemaining)
	} else if p.State == RUNNING_BREAK {
		txt = IconPlay + "  " + fmtDuration(p.BreakRemaining)
	}

	return &Block{FullText: txt, Color: ColorLime}
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}

func NewPomodoroAddon() *Addon {
	interval := 1000 * time.Millisecond
	return &Addon{
		UpdateInterval: interval,
		Updater:        &pomodoro{State: RUNNING_TASK, UpdateInterval: interval},
	}
}
