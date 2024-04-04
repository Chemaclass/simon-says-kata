package main_test

import (
	"testing"

	. "github.com/JesusValera/simon-says-kata"
)

func TestWIPGame(t *testing.T) {
	game := Game{}
	game.Play()
	ok := game.UserInput("y")

	if !ok {
		t.Error("error")
	}
}
