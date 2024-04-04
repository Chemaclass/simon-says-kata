package main

type Game struct {
	list []string
}

func (g *Game) Play() {

}

func (g *Game) UserInput(input string) bool {
	return true
}
