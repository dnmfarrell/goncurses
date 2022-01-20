// ncurses_example.go
package goncurses_test

import (
	"testing"

	"github.com/dnmfarrell/goncurses"
)

func TestInit(t *testing.T) {
	_, err := goncurses.Init()
	if err != nil {
		t.Fatal(err)
	}
	defer goncurses.End()
}

func TestGetChar(t *testing.T) {
	_, err := goncurses.Init()
	if err != nil {
		t.Fatal(err)
	}
	defer goncurses.End()
	var in goncurses.Char = 'n'
	goncurses.UnGetChar(in)
	out := goncurses.GetChar()
	if in != out {
		t.Errorf("Expected %d but got %d", in, out)
	}
}

func TestGetWChar(t *testing.T) {
	_, err := goncurses.Init()
	if err != nil {
		t.Fatal(err)
	}
	defer goncurses.End()
	var in goncurses.WChar = 'ń'
	err = goncurses.UnGetWChar(in)
	if err != nil {
		t.Errorf("UngetWChar returns an error: %s", err)
	}
	out := goncurses.GetWChar()
	if in != out {
		t.Errorf("Expected %d but got %d", in, out)
	}
}

func TestAddChar(t *testing.T) {
	_, err := goncurses.Init()
	if err != nil {
		t.Fatal(err)
	}
	defer goncurses.End()
	var in goncurses.Char = 'n'
	err = goncurses.AddChar(in)
	if err != nil {
		t.Errorf("AddChar returns an error: %s", err)
	}
}

func TestAddWChar(t *testing.T) {
	_, err := goncurses.Init()
	if err != nil {
		t.Fatal(err)
	}
	defer goncurses.End()
	var in goncurses.WChar = 'ń'
	err = goncurses.AddWChar(in)
	if err != nil {
		t.Errorf("UngetWChar returns an error: %s", err)
	}
}

