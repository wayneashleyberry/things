package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/wayneashleyberry/things/pkg/url"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "show":
		switch os.Args[2] {
		case "inbox":
			showInbox()
		case "today":
			showToday()
		case "upcoming":
			showUpcoming()
		case "anytime":
			showAnytime()
		case "someday":
			showSomeday()
		case "logbook":
			showLogbook()
		case "trash":
			showTrash()
		}
		break
	default:
		log.Fatal("invalid command")
	}
}

func open(url string) {
	cmd := exec.Command("open", url)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func showInbox() {
	open(url.ShowInbox)
}

func showToday() {
	open(url.ShowToday)
}

func showAnytime() {
	open(url.ShowAnytime)
}

func showUpcoming() {
	open(url.ShowUpcoming)
}

func showSomeday() {
	open(url.ShowSomeday)
}

func showLogbook() {
	open(url.ShowLogbook)
}

func showTrash() {
	open(url.ShowTrash)
}

func usage() {
	fmt.Println(`
Things Unofficial CLI

Usage:

	things command [arguments]

The commands are:

	show	Show a list in the app
	`)
}
