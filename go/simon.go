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
	if n == 3 {
		g.list = append(g.list, "red")
	} else if n == 2 {
		g.list = append(g.list, "blue")
	} else if n == 1 {
		g.list = append(g.list, "yellow")
	} else {
		g.list = append(g.list, "green")
	}
	return g.list
}

func (g *Game) UserInput(input string) bool {
	return true
}
