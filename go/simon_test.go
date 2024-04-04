package main_test

import (
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

func areEqual(slice1, slice2 []string) bool {
	// If the lengths of the slices are different, they can't be equal
	if len(slice1) != len(slice2) {
		return false
	}

	// Iterate over each element of the slices and compare them
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	// If all elements are equal, return true
	return true
}

func TestColorsAfterPlay(t *testing.T) {
	game := Game{}
	colors := game.Play()

	if !areEqual(colors, []string{"yellow"}) {
		t.Error("error")
	}
}
