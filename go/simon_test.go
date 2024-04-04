package main_test

import (
	"reflect"
	"testing"

	. "github.com/JesusValera/simon-says-kata"
)

type MockRandomGenerator struct {
	values []string
	index  int
}

func NewMockRandomGenerator(values []string) *MockRandomGenerator {
	return &MockRandomGenerator{
		values: values,
		index:  0,
	}
}

func (rg *MockRandomGenerator) Generate() string {
	value := rg.values[rg.index]
	rg.index++
	if rg.index >= len(rg.values) {
		rg.index = 0
	}
	return value
}

func TestPlayReturnsYellow(t *testing.T) {
	game := NewGame(NewMockRandomGenerator([]string{"yellow"}))
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"yellow"}) {
		t.Error("error")
	}
}

func TestPlayReturnsRed(t *testing.T) {
	game := NewGame(NewMockRandomGenerator([]string{"red"}))
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"red"}) {
		t.Error("error")
	}
}

func TestPlayReturnsBlue(t *testing.T) {
	game := NewGame(NewMockRandomGenerator([]string{"blue"}))
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"blue"}) {
		t.Error("error")
	}
}

func TestPlayReturnsGreen(t *testing.T) {
	game := NewGame(NewMockRandomGenerator([]string{"green"}))
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"green"}) {
		t.Errorf("errors %s", colors[0])
	}
}

func TestPlayReturnsTwoColors(t *testing.T) {
	game := NewGame(NewMockRandomGenerator([]string{"green", "red"}))
	game.Play()
	colors := game.Play()

	if !reflect.DeepEqual(colors, []string{"green", "red"}) {
		t.Errorf("errors %s", colors)
	}
}

func TestValidUserInput(t *testing.T) {
	game := NewGame(NewMockRandomGenerator([]string{"green"}))
	game.Play()
	ok := game.UserInput("g")

	if !ok {
		t.Error("error")
	}
}

func TestInvalidUserInput(t *testing.T) {
	game := NewGame(NewMockRandomGenerator([]string{"yellow"}))
	game.Play()
	ok := game.UserInput("r")

	if ok {
		t.Error("error")
	}
}
