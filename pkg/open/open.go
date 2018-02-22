package open

import (
	"log"
	"os/exec"
)

func Open(url string) {
	cmd := exec.Command("open", url)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
