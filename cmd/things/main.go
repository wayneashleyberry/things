package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wayneashleyberry/things/pkg/open"
	"github.com/wayneashleyberry/things/pkg/url"
)

func main() {
	completed := false

	var cmdAdd = &cobra.Command{
		Use:   "add",
		Short: "Add a new to-do",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			add(args[0], completed)
		},
	}

	cmdAdd.Flags().BoolVarP(&completed, "completed", "", false, "Boolean. Whether or not the to-do should be set to complete. Default: false. Ignored if canceled is also set to true.")

	var cmdAddJSON = &cobra.Command{
		Use:   "add-json",
		Short: "Add a new to-do from json",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			addJSON(args[0])
		},
	}

	var cmdAddProject = &cobra.Command{
		Use:   "add-project",
		Short: "Add a new project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			addProject(args[0])
		},
	}

	var cmdShow = &cobra.Command{
		Use:   "show",
		Short: "Show a specific list",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}

	var cmdSearch = &cobra.Command{
		Use:   "search",
		Short: "Search in everything",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			search(args[0])
		},
	}

	var cmdShowInbox = &cobra.Command{
		Use:   "inbox",
		Short: "Show the inbox",
		Run: func(cmd *cobra.Command, args []string) {
			showInbox()
		},
	}

	var cmdShowToday = &cobra.Command{
		Use:   "today",
		Short: "Show the today smart list",
		Run: func(cmd *cobra.Command, args []string) {
			showToday()
		},
	}

	var cmdShowAnytime = &cobra.Command{
		Use:   "anytime",
		Short: "Show the anytime smart list",
		Run: func(cmd *cobra.Command, args []string) {
			showAnytime()
		},
	}

	var cmdShowUpcoming = &cobra.Command{
		Use:   "upcoming",
		Short: "Show the upcoming smart list",
		Run: func(cmd *cobra.Command, args []string) {
			showUpcoming()
		},
	}

	var cmdShowSomeday = &cobra.Command{
		Use:   "someday",
		Short: "Show the someday smart list",
		Run: func(cmd *cobra.Command, args []string) {
			showSomeday()
		},
	}

	var cmdShowLogbook = &cobra.Command{
		Use:   "logbook",
		Short: "Show the logbook",
		Run: func(cmd *cobra.Command, args []string) {
			showLogbook()
		},
	}

	var cmdShowTrash = &cobra.Command{
		Use:   "trash",
		Short: "Show trash",
		Run: func(cmd *cobra.Command, args []string) {
			showTrash()
		},
	}

	var cmdShowTask = &cobra.Command{
		Use:   "task [id]",
		Short: "Show a specific task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			showTask(args[0])
		},
	}

	var cmdShowQuery = &cobra.Command{
		Use:   "query [name]",
		Short: "Show an area, project, tag or a built-in list by name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			showQuery(args[0])
		},
	}

	var rootCmd = &cobra.Command{Use: "things"}
	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdAddJSON)
	rootCmd.AddCommand(cmdAddProject)
	rootCmd.AddCommand(cmdShow)
	rootCmd.AddCommand(cmdSearch)
	cmdShow.AddCommand(cmdShowInbox)
	cmdShow.AddCommand(cmdShowToday)
	cmdShow.AddCommand(cmdShowAnytime)
	cmdShow.AddCommand(cmdShowUpcoming)
	cmdShow.AddCommand(cmdShowSomeday)
	cmdShow.AddCommand(cmdShowLogbook)
	cmdShow.AddCommand(cmdShowTrash)
	cmdShow.AddCommand(cmdShowTask)
	cmdShow.AddCommand(cmdShowQuery)
	rootCmd.Execute()
}

func add(title string, completed bool) {
	a := url.Add{
		Title:     title,
		Completed: completed,
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
