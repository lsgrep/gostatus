package addon

import (
	"fmt"
	"os"
	"os/exec"
)

func Notify(msg string) {
	c := exec.Command("sh", "-c", fmt.Sprintf("notify-send '%s'", msg))
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	if err := c.Run(); err != nil {
		logger.Error(err)
	}
}
