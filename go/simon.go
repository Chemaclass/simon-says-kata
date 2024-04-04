package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var possibleColors = [...]string{
	"red",
	"green",
	"blue",
	"yellow",
}

type RandomGenerator interface {
	Generate() string
}

type RealRandomGenerator struct{}

func (rg *RealRandomGenerator) Generate() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return possibleColors[rand.Int()%4]
}

type Game struct {
	list            []string
	randomGenerator RandomGenerator
}

func NewGame(randomGenerator RandomGenerator) *Game {
	return &Game{
		randomGenerator: randomGenerator,
	}
}

func (g *Game) Play() []string {
	w := g.randomGenerator.Generate()
	g.list = append(g.list, w)
	return g.list
}

func (g *Game) UserInput(input string) bool {
	if input[0] == g.list[0][0] {
		return true
	}
	return false
}

func main() {
	game := NewGame(&RealRandomGenerator{})
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Printf("Simon says: %s\n", strings.Join(game.Play(), " "))
		fmt.Print("Your turn: ")
		text, _ := reader.ReadString('\n')

		if !game.UserInput(text) {
			fmt.Println("Game Over!")
			break
		}
		fmt.Println("Good job!")
	}
}
