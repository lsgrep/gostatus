package addon

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// Mute: no
var mutedRegexStr = `.*Mute:\s(\S+)`

// Volume: front-left: 33835 /  52% / -17.23 dB,   front-right: 33835 /  52% / -17.23 dB
var volumeRegexStr = `(\d+%)`

var mutedReg *regexp.Regexp
var volumeReg *regexp.Regexp

func init() {
	mutedReg = regexp.MustCompile(mutedRegexStr)
	volumeReg = regexp.MustCompile(volumeRegexStr)
}

// TODO replace this with any elegant solution
func GetVolume() (bool, string) {
	buf := bytes.NewBufferString("")
	c := exec.Command("sh", "-c", "pactl list sinks | grep -A 20 \"RUNNING\"")
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

func (vs *volumeStatus) Update() *Block {
	muted, v := GetVolume()
	var ret string
	if muted {
		ret = "muted"
	} else {
		ret = v
	}
	fullTxt := fmt.Sprintf(" %s  %s", IconVolume, ret)
	return &Block{FullText: fullTxt}
}

func NewVolumeAddon() *Addon {
	v := &volumeStatus{}
	aa := Addon{
		UpdateInterval: 1000 * time.Millisecond,
		Updater:        v}
	return &aa
}
