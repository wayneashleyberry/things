package url

import (
	gourl "net/url"
	"time"
)

const Scheme = "things:///"

const ShowInbox = Scheme + "show?id=inbox"
const ShowToday = Scheme + "show?id=today"
const ShowUpcoming = Scheme + "show?id=upcoming"
const ShowAnytime = Scheme + "show?id=anytime"
const ShowSomeday = Scheme + "show?id=someday"

const ShowLogbook = Scheme + "show?id=logbook"
const ShowTrash = Scheme + "show?id=trash"

const ShowID = Scheme + "show?id=%s"
const ShowQuery = Scheme + "show?query=%s"

const Search = Scheme + "search?query=%s"

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

func (a Add) URL() string {
	v := gourl.Values{}
	v.Add("title", a.Title)
	return Scheme + "add?" + v.Encode()
}

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

func (a AddProject) URL() string {
	v := gourl.Values{}
	v.Add("title", a.Title)
	return Scheme + "add-project?" + v.Encode()
}

type AddJSON struct {
	JSON string
}

func (a AddJSON) URL() string {
	v := gourl.Values{}
	v.Add("data", a.JSON)
	return Scheme + "add-json?" + v.Encode()
}
