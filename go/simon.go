package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
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

func main() {
	game := NewGame(1)
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
