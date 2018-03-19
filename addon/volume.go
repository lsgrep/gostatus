package addon

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Mute: no
var mutedRegexStr = `.*Mute:\s(\S+)`

// Volume: front-left: 33835 /  52% / -17.23 dB,   front-right: 33835 /  52% / -17.23 dB
var volumeRegexStr = `(\d+%)`

var mutedReg *regexp.Regexp
var volumeReg *regexp.Regexp

func init() {
	var err error
	mutedReg, err = regexp.Compile(mutedRegexStr)
	if err != nil {
		panic(err)
	}

	volumeReg, err = regexp.Compile(volumeRegexStr)
	if err != nil {
		panic(err)
	}
}

// TODO replace this with any elegant solution
func GetVolume() (bool, string) {
	buf := bytes.NewBufferString("")
	c := exec.Command("sh", "-c", "pactl list sinks | grep -B 10000 \"SUSPENDED\"")
	c.Stderr = os.Stderr
	c.Stdout = buf
	if err := c.Run(); err != nil {
		// TODO log error
		return false, ""
	}
	output := buf.String()
	lines := strings.Split(output, "\n")

	var muted bool
	var volume string

	for _, line := range lines {
		if strings.Contains(line, "Mute:") {
			matches := mutedReg.FindStringSubmatch(line)
			// TODO `no` is locale specific, so it might fail with German etc
			if len(matches) > 1 && strings.ToLower(mutedReg.FindStringSubmatch(line)[1]) == "no" {
				muted = false
			} else {
				muted = true
			}
		}

		if strings.Contains(line, "Volume: front-left:") {
			matches := volumeReg.FindStringSubmatch(line)
			if len(matches) > 1 {
				volume = matches[0]
			}
		}
	}
	return muted, volume
}

type volumeStatus struct {
}

func (vs *volumeStatus) Update() string {
	muted, v := GetVolume()
	if muted {
		return "muted"
	}
	return v
}

func NewVolumeAddon() *Addon {
	v := &volumeStatus{}
	aa := Addon{
		Icon:             "\uf028",
		UpdateIntervalMs: 1000,
		Updater:          v}
	return &aa
}
