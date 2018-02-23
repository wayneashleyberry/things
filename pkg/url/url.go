package url

import (
	neturl "net/url"
	"strings"
	"time"
)

// Scheme is the base URL scheme used by all other URL's
const Scheme = "things:///"

// ShowInbox is the URL used to show the Inbox
const ShowInbox = Scheme + "show?id=inbox"

// ShowToday is the URL used to show the Today smart list
const ShowToday = Scheme + "show?id=today"

// ShowUpcoming is the URL used to show the Upcoming smart list
const ShowUpcoming = Scheme + "show?id=upcoming"

// ShowAnytime is the URL used to show the Anytime smart list
const ShowAnytime = Scheme + "show?id=anytime"

// ShowSomeday is the URL used to show the Someday smart list
const ShowSomeday = Scheme + "show?id=someday"

// ShowLogbook is the URL for showing the logbook
const ShowLogbook = Scheme + "show?id=logbook"

// ShowTrash is the URL for showing the trash
const ShowTrash = Scheme + "show?id=trash"

// ShowID is the URL for showing a specific to-do by ID
const ShowID = Scheme + "show?id=%s"

// ShowQuery is the URL for showing a area, project, tag or a
// built-in list by name
const ShowQuery = Scheme + "show?query=%s"

// Search is the URL for opening a search query
const Search = Scheme + "search?query=%s"

// Add contains parameters for adding a new to-do
type Add struct {
	Title          string
	Completed      bool
	Canceled       bool
	ShowQuickEntry bool
	Reveal         bool
	Notes          string
	ChecklistItems []string
	When           time.Time
	Deadline       time.Time
	Tags           []string
	List           string
	ListID         string
	Heading        string
}

// URL builds the URL
func (a Add) URL() string {
	v := neturl.Values{}
	v.Add("title", a.Title)
	if a.Completed {
		v.Add("completed", "true")
	}
	if a.Canceled {
		v.Add("canceled", "true")
	}
	if a.Reveal {
		v.Add("reveal", "true")
	}
	if a.ShowQuickEntry {
		v.Add("show-quick-entry", "true")
	}
	if a.Notes != "" {
		v.Add("notes", a.Notes)
	}
	url := Scheme + "add?" + v.Encode()
	return strings.TrimRight(url, "&")
}

// AddProject contains parameters for adding a new project
type AddProject struct {
	Title     string
	Completed bool
	Canceled  bool
	Reveal    bool
	Notes     string
	When      time.Time
	Deadline  time.Time
	Tags      []string
	Area      string
	AreaID    string
	ToDos     []string
}

// URL builds the URL
func (a AddProject) URL() string {
	v := neturl.Values{}
	v.Add("title", a.Title)
	return Scheme + "add-project?" + v.Encode()
}

// AddJSON contains parameters for adding a new to-do from json
type AddJSON struct {
	JSON string
}

// URL builds the URL
func (a AddJSON) URL() string {
	v := neturl.Values{}
	v.Add("data", a.JSON)
	return Scheme + "add-project?" + v.Encode()
}
