package url

import (
	neturl "net/url"
	"strings"
)

// Scheme is the base URL scheme used by all other URL's
const Scheme = "things:///"

// Show is the URL used to show lists
type Show struct {
	ID      string
	Query   string
	Filters []string
}

// URL implementation
func (s Show) URL() string {
	v := neturl.Values{}
	if s.ID != "" {
		v.Add("id", s.ID)
	}
	if s.Query != "" {
		v.Add("query", s.Query)
	}
	if len(s.Filters) > 0 {
		v.Add("filter", strings.Join(s.Filters, ","))
	}
	return Scheme + "show?" + v.Encode()
}

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
	When           string
	Deadline       string
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
	if len(a.ChecklistItems) > 0 {
		v.Add("checklist-items", strings.Join(a.ChecklistItems, "\n"))
	}
	if a.When != "" {
		v.Add("when", a.When)
	}
	if a.Deadline != "" {
		v.Add("deadline", a.Deadline)
	}
	if len(a.Tags) > 0 {
		v.Add("tags", strings.Join(a.Tags, ","))
	}
	if a.List != "" {
		v.Add("list", a.List)
	}
	if a.ListID != "" {
		v.Add("list-id", a.ListID)
	}
	if a.Heading != "" {
		v.Add("heading", a.Heading)
	}
	return Scheme + "add?" + strings.Replace(v.Encode(), "+", "%20", -1)
}

// AddProject contains parameters for adding a new project
type AddProject struct {
	Title     string
	Completed bool
	Canceled  bool
	Reveal    bool
	Notes     string
	When      string
	Deadline  string
	Tags      []string
	Area      string
	AreaID    string
	ToDos     []string
}

// URL builds the URL
func (a AddProject) URL() string {
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
	if a.Notes != "" {
		v.Add("notes", a.Notes)
	}
	if a.When != "" {
		v.Add("when", a.When)
	}
	if a.Deadline != "" {
		v.Add("deadline", a.Deadline)
	}
	if len(a.Tags) > 0 {
		v.Add("tags", strings.Join(a.Tags, ","))
	}
	if a.Area != "" {
		v.Add("area", a.Area)
	}
	if a.AreaID != "" {
		v.Add("area-id", a.AreaID)
	}
	if len(a.ToDos) > 0 {
		v.Add("to-dos", strings.Join(a.ToDos, "\n"))
	}
	return Scheme + "add-project?" + strings.Replace(v.Encode(), "+", "%20", -1)
}

// AddJSON contains parameters for adding a new to-do from json
type AddJSON struct {
	JSON   string
	Reveal bool
}

// URL builds the URL
func (a AddJSON) URL() string {
	v := neturl.Values{}
	v.Add("data", a.JSON)
	if a.Reveal {
		v.Add("reveal", "true")
	}
	return Scheme + "add-project?" + v.Encode()
}
