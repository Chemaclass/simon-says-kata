package main

import (
	"math/rand"
)

type Game struct {
	list []string
	rand *rand.Rand
}

func NewGame(seed int) Game {
	game := Game{}
	game.rand = rand.New(rand.NewSource(int64(seed)))
	return game
}

func (g *Game) Play() []string {
	n := g.rand.Int() % 4
	if n == 2 {
		return []string{"blue"}
	} else if n == 1 {
		return []string{"yellow"}
	} else {
		return []string{"red"}
	}

}

func (g *Game) UserInput(input string) bool {
	return true
}
