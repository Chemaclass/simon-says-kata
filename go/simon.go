package main

type Game struct {
	list []string
}

func (g *Game) Play() []string {
	return []string{"yellow"}
}

func (g *Game) UserInput(input string) bool {
	return true
}
