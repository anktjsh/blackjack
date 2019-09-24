package gameutil

import (
	"math/rand"
	"time"
)

type Deck struct {
	cards []PlayingCard
}

func (d *Deck) Create() *Deck {
	var cards []PlayingCard
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			cards = append(cards, PlayingCard{Suit(i), Card(j + 1)})
		}
	}
	d.cards = cards
	d.Shuffle()
	return d
}

func (d *Deck) Size() int {
	return len(d.cards)
}

func (d *Deck) PopCard() PlayingCard {
	result := d.cards[0]
	d.cards = d.cards[1:]
	return result
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}
