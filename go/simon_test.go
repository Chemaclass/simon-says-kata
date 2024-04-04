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

func TestColorsAfterPlay(t *testing.T) {
	game := NewGame(3)
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"yellow"}) {
		t.Error("error")
	}
}

func TestPlayReturnsOtherColors(t *testing.T) {
	game := NewGame(2)
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"red"}) {
		t.Error("error")
	}
}
