package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wayneashleyberry/things/pkg/open"
	"github.com/wayneashleyberry/things/pkg/url"
)

func handleURL(url string, printURL bool) {
	if printURL {
		fmt.Println(url)
	} else {
		open.Open(url)
	}
}

func main() {
	printURL := false
	completed := false
	canceled := false
	showQuickEntry := false
	reveal := false
	notes := ""
	checklistItems := []string{}
	when := ""
	deadline := ""
	tags := []string{}
	list := ""
	listID := ""
	heading := ""

	var cmdAdd = &cobra.Command{
		Use:   "add [title]",
		Short: "Add a new to-do",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			a := url.Add{
				Title:          args[0],
				Completed:      completed,
				Canceled:       canceled,
				ShowQuickEntry: showQuickEntry,
				Reveal:         reveal,
				Notes:          notes,
				ChecklistItems: checklistItems,
				When:           when,
				Deadline:       deadline,
				Tags:           tags,
				List:           list,
				ListID:         listID,
				Heading:        heading,
			}
			handleURL(a.URL(), printURL)
		},
	}

	cmdAdd.Flags().BoolVarP(&completed, "completed", "", false, "Whether or not the to-do should be set to complete. Default: false. Ignored if canceled is also set to true.")
	cmdAdd.Flags().BoolVarP(&canceled, "canceled", "", false, "Whether or not the to-do should be set to canceled. Default: false. Takes priority over completed.")
	cmdAdd.Flags().BoolVarP(&showQuickEntry, "show-quick-entry", "", false, "Whether or not to show the quick entry dialog (populated with the provided data) instead of adding a new to-do. Ignored if titles is specified. Default: false.")
	cmdAdd.Flags().BoolVarP(&reveal, "reveal", "", true, "Whether or not to navigate to and show the newly created to-do. If multiple to-dos have been created, the first one will be shown. Ignored if show-quick-entry is also set to true. Default: false.")
	cmdAdd.Flags().StringVarP(&notes, "notes", "", "", "The text to use for the notes field of the to-do. Maximum unencoded length: 10,000 characters.")
	cmdAdd.Flags().StringArrayVarP(&checklistItems, "checklist-item", "", []string{}, "Checklist items to add to the to-do (maximum of 100).")
	cmdAdd.Flags().StringArrayVarP(&tags, "tag", "", []string{}, "Strings corresponding to the titles of tags. Does not apply a tag if the specified tag doesn’t exist.")
	cmdAdd.Flags().StringVarP(&when, "when", "", "", "Possible values: today, tomorrow, evening, anytime, someday, a date string, or a date time string. Using a date time string adds a reminder for that time. The time component is ignored if anytime or someday is specified.")
	cmdAdd.Flags().StringVarP(&deadline, "deadline", "", "", "The deadline to apply to the to-do.")
	cmdAdd.Flags().StringVarP(&list, "list", "", "", "The title of a project or area to add to. Ignored if list-id is present.")
	cmdAdd.Flags().StringVarP(&listID, "list-id", "", "", "The ID of a project or area to add to. Takes precedence over list.")
	cmdAdd.Flags().StringVarP(&heading, "heading", "", "", "The title of a heading within a project to add to. Ignored if a project is not specified, or if the heading doesn't exist.")

	var cmdAddJSON = &cobra.Command{
		Use:   "add-json [json]",
		Short: "Add a new to-do from json",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			addJSON(args[0], reveal, printURL)
		},
	}

	cmdAddJSON.Flags().BoolVarP(&reveal, "reveal", "", true, "Whether or not to navigate to and show the newly created to-do. If multiple to-dos have been created, the first one will be shown. Ignored if show-quick-entry is also set to true. Default: false.")

	area := ""
	areaID := ""
	todos := []string{}

	var cmdAddProject = &cobra.Command{
		Use:   "add-project [title]",
		Short: "Add a new project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			a := url.AddProject{
				Title:     args[0],
				Completed: completed,
				Canceled:  canceled,
				Reveal:    reveal,
				Notes:     notes,
				When:      when,
				Deadline:  deadline,
				Tags:      tags,
				Area:      area,
				AreaID:    areaID,
				ToDos:     todos,
			}
			handleURL(a.URL(), printURL)
		},
	}

	cmdAddProject.Flags().BoolVarP(&completed, "completed", "", false, "Whether or not the to-do should be set to complete. Default: false. Ignored if canceled is also set to true.")
	cmdAddProject.Flags().BoolVarP(&canceled, "canceled", "", false, "Whether or not the to-do should be set to canceled. Default: false. Takes priority over completed.")
	cmdAddProject.Flags().BoolVarP(&reveal, "reveal", "", true, "Whether or not to navigate to and show the newly created to-do. If multiple to-dos have been created, the first one will be shown. Ignored if show-quick-entry is also set to true. Default: false.")
	cmdAddProject.Flags().StringVarP(&notes, "notes", "", "", "The text to use for the notes field of the to-do. Maximum unencoded length: 10,000 characters.")
	cmdAddProject.Flags().StringVarP(&when, "when", "", "", "Possible values: today, tomorrow, evening, anytime, someday, a date string, or a date time string. Using a date time string adds a reminder for that time. The time component is ignored if anytime or someday is specified.")
	cmdAddProject.Flags().StringVarP(&deadline, "deadline", "", "", "The deadline to apply to the to-do.")
	cmdAddProject.Flags().StringArrayVarP(&tags, "tag", "", []string{}, "Strings corresponding to the titles of tags. Does not apply a tag if the specified tag doesn’t exist.")
	cmdAddProject.Flags().StringVarP(&area, "area", "", "", "The title of an area to add to. Ignored if area-id is present.")
	cmdAddProject.Flags().StringVarP(&areaID, "area-id", "", "", "The ID of an area to add to. Takes precedence over area.")
	cmdAddProject.Flags().StringArrayVarP(&todos, "to-do", "", []string{}, "Titles of to-dos to create inside the project.")

	var cmdSearch = &cobra.Command{
		Use:   "search [query]",
		Short: "Search in everything",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			search(args[0], printURL)
		},
	}

	filters := []string{}

	var cmdShow = &cobra.Command{
		Use:   "show",
		Short: "Show a specific list",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			showQuery(filters, args[0], printURL)
		},
	}

	cmdShow.PersistentFlags().StringArrayVarP(&filters, "filter", "", []string{}, "Strings corresponding to the titles of tags that the list should be filtered by.")

	var cmdShowInbox = &cobra.Command{
		Use:   "inbox",
		Short: "Show the inbox",
		Run: func(cmd *cobra.Command, args []string) {
			showInbox(filters, printURL)
		},
	}

	var cmdShowToday = &cobra.Command{
		Use:   "today",
		Short: "Show the today smart list",
		Run: func(cmd *cobra.Command, args []string) {
			showToday(filters, printURL)
		},
	}

	var cmdShowAnytime = &cobra.Command{
		Use:   "anytime",
		Short: "Show the anytime smart list",
		Run: func(cmd *cobra.Command, args []string) {
			showAnytime(filters, printURL)
		},
	}

	var cmdShowUpcoming = &cobra.Command{
		Use:   "upcoming",
		Short: "Show the upcoming smart list",
		Run: func(cmd *cobra.Command, args []string) {
			showUpcoming(filters, printURL)
		},
	}

	var cmdShowSomeday = &cobra.Command{
		Use:   "someday",
		Short: "Show the someday smart list",
		Run: func(cmd *cobra.Command, args []string) {
			showSomeday(filters, printURL)
		},
	}

	var cmdShowLogbook = &cobra.Command{
		Use:   "logbook",
		Short: "Show the logbook",
		Run: func(cmd *cobra.Command, args []string) {
			showLogbook(filters, printURL)
		},
	}

	var cmdShowTrash = &cobra.Command{
		Use:   "trash",
		Short: "Show trash",
		Run: func(cmd *cobra.Command, args []string) {
			showTrash(filters, printURL)
		},
	}

	var cmdShowTask = &cobra.Command{
		Use:   "task [id]",
		Short: "Show a specific task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			showTask(filters, args[0], printURL)
		},
	}

	var cmdShowQuery = &cobra.Command{
		Use:   "query [name]",
		Short: "Show an area, project, tag or a built-in list by name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			showQuery(filters, args[0], printURL)
		},
	}

	var cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Print things version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("1.0.5")
		},
	}

	var rootCmd = &cobra.Command{
		Use: "things",
	}

	rootCmd.PersistentFlags().BoolVar(&printURL, "print", false, "Print the URL instead of opening it")

	rootCmd.AddCommand(cmdVersion)
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

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addJSON(json string, reveal, printURL bool) {
	a := url.AddJSON{
		JSON:   json,
		Reveal: reveal,
	}
	handleURL(a.URL(), printURL)
}

func search(query string, printURL bool) {
	handleURL(fmt.Sprintf(url.Search, query), printURL)
}

func showInbox(filters []string, printURL bool) {
	s := url.Show{
		ID:      url.ShowInboxID,
		Filters: filters,
	}
	handleURL(s.URL(), printURL)

}

func showToday(filters []string, printURL bool) {
	s := url.Show{
		ID:      url.ShowTodayID,
		Filters: filters,
	}
	handleURL(s.URL(), printURL)
}

func showAnytime(filters []string, printURL bool) {
	s := url.Show{
		ID:      url.ShowAnytimeID,
		Filters: filters,
	}
	handleURL(s.URL(), printURL)
}

func showUpcoming(filters []string, printURL bool) {
	s := url.Show{
		ID:      url.ShowUpcomingID,
		Filters: filters,
	}
	handleURL(s.URL(), printURL)
}

func showSomeday(filters []string, printURL bool) {
	s := url.Show{
		ID:      url.ShowSomedayID,
		Filters: filters,
	}
	handleURL(s.URL(), printURL)
}

func showLogbook(filters []string, printURL bool) {
	s := url.Show{
		ID:      url.ShowLogbookID,
		Filters: filters,
	}
	handleURL(s.URL(), printURL)
}

func showTrash(filters []string, printURL bool) {
	s := url.Show{
		ID:      url.ShowTrashID,
		Filters: filters,
	}
	handleURL(s.URL(), printURL)
}

func showTask(filters []string, id string, printURL bool) {
	s := url.Show{
		ID:      id,
		Filters: filters,
	}
	handleURL(s.URL(), printURL)
}

func showQuery(filters []string, query string, printURL bool) {
	s := url.Show{
		Query:   query,
		Filters: filters,
	}
	handleURL(s.URL(), printURL)
}
