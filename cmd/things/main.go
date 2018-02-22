package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wayneashleyberry/things/pkg/open"
	"github.com/wayneashleyberry/things/pkg/url"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "add":
		add(os.Args[2])
	case "add-project":
		addProject(os.Args[2])
	case "add-json":
		addJSON(os.Args[2])
	case "search":
		search(os.Args[2])
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
		case "task":
			showTask(os.Args[3])
		case "query":
			showQuery(os.Args[3])
		}
		break
	case "help":
		switch os.Args[2] {
		case "show":
			usageShow()
		case "add":
			usageAdd()
		case "add-project":
			usageAddProject()
		default:
			usage()
		}

	default:
		log.Fatal("invalid command")
	}
}

func add(title string) {
	a := url.Add{
		Title: title,
	}
	open.Open(a.URL())
}

func addJSON(json string) {
	a := url.AddJSON{
		JSON: json,
	}
	open.Open(a.URL())
}

func addProject(title string) {
	a := url.AddProject{
		Title: title,
	}
	open.Open(a.URL())
}

func search(query string) {
	open.Open(fmt.Sprintf(url.Search, query))
}

func showInbox() {
	open.Open(url.ShowInbox)
}

func showToday() {
	open.Open(url.ShowToday)
}

func showAnytime() {
	open.Open(url.ShowAnytime)
}

func showUpcoming() {
	open.Open(url.ShowUpcoming)
}

func showSomeday() {
	open.Open(url.ShowSomeday)
}

func showLogbook() {
	open.Open(url.ShowLogbook)
}

func showTrash() {
	open.Open(url.ShowTrash)
}

func showTask(id string) {
	open.Open(fmt.Sprintf(url.ShowID, id))
}

func showQuery(query string) {
	open.Open(fmt.Sprintf(url.ShowQuery, query))
}

func usage() {
	fmt.Println(`
Unofficial CLI for Things

Usage:

	things command [arguments]

The commands are:

	add		Add a new task
	add-json	Add a new task using json
	add-project 	Add a new project
	show		Show a list in the app
	search		Search across everything

Use "things help [command]" for more information about a command.
	`)
}

func usageShow() {
	fmt.Println(`
Usage:

	things show argument

The supported arguments are:

	inbox		Show the Inbox
	today		Show the Today list
	upcoming	Show the Upcoming list
	anytimg		Show the Anytime list
	someday		Show the Someday list
	logboox		Show the Logbook
	trash		Show the Trash
	id		Show a task by ID
	query		Show a list by query
`)
}

func usageAdd() {
	fmt.Println(`
Usage:

	things add title
`)
}

func usageAddProject() {
	fmt.Println(`
Usage:

	things add-project title
`)
}
