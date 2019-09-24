package gameutil

import "bufio"

type Game struct {
	deck   *Deck
	money  int
	reader *bufio.Reader
}

func (g *Game) Start() {

}
