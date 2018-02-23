package open

import (
	"log"
	"os/exec"
)

// Open wraps the shell command "open", considering Things is only available
// on macOS this should probably be fine.
func Open(url string) {
	cmd := exec.Command("open", url)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
