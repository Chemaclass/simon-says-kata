package main_test

import (
	"reflect"
	"testing"

	. "github.com/JesusValera/simon-says-kata"
)

/*
func TestWIPGame(t *testing.T) {
	game := Game{}
	game.Play()
	ok := game.UserInput("y")

	if !ok {
		t.Error("error")
	}
}
*/

func TestPlayReturnsYellow(t *testing.T) {
	game := NewGame(3)
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"yellow"}) {
		t.Error("error")
	}
}

func TestPlayReturnsRed(t *testing.T) {
	game := NewGame(2)
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"red"}) {
		t.Error("error")
	}
}

func TestPlayReturnsBlue(t *testing.T) {
	game := NewGame(1)
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"blue"}) {
		t.Error("error")
	}
}

func TestPlayReturnsGreen(t *testing.T) {
	game := NewGame(5)
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"green"}) {
		t.Errorf("errors %s", colors[0])
	}
}

func TestPlayReturnsTwoColors(t *testing.T) {
	game := NewGame(5)
	game.Play()
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"green", "red"}) {
		t.Errorf("errors %s", colors[0])
	}
}

func TestValidUserInput(t *testing.T) {
	game := NewGame(5)
	game.Play()
	ok := game.UserInput("g")

	if !ok {
		t.Error("error")
	}
}

func TestInvalidUserInput(t *testing.T) {
	game := NewGame(5)
	game.Play()
	ok := game.UserInput("r")

	if ok {
		t.Error("error")
	}
}
