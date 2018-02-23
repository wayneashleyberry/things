package url

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := Add{
		Title: "Hello, World!",
	}
	exptected := "things:///add?title=Hello%2C%20World%21"
	got := a.URL()
	if got != exptected {
		t.Errorf("expected %s, got %s", exptected, got)
	}
}

func TestProject(t *testing.T) {
	a := AddProject{
		Title: "Hello, World!",
	}
	exptected := "things:///add-project?title=Hello%2C%20World%21"
	got := a.URL()
	if got != exptected {
		t.Errorf("expected %s, got %s", exptected, got)
	}
}

func TestShowQuery(t *testing.T) {
	s := Show{
		Query: "test project",
	}
	exptected := "things:///show?query=test%20project"
	got := s.URL()
	if got != exptected {
		t.Errorf("expected %s, got %s", exptected, got)
	}
}

func TestShowQueryWithSpaces(t *testing.T) {
	s := Show{
		Query: "+test",
	}
	exptected := "things:///show?query=%2Btest"
	got := s.URL()
	if got != exptected {
		t.Errorf("expected %s, got %s", exptected, got)
	}
}
