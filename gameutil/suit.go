package gameutil

import "github.com/gookit/color"

var suits = [4]string{"Hearts", "Clubs", "Diamonds", "Spades"}
var visualSuits = [4]string{"♥", "♣", "♦", "♠"}

type Suit int

const (
	Hearts   Suit = 0
	Clubs    Suit = 1
	Diamonds Suit = 2
	Spades   Suit = 3
)

func (suit Suit) String() string {
	return suits[suit]
}

func (suit Suit) Visual() string {
	return visualSuits[suit]
}

func (suit Suit) Color() color.Color {
	if suit == Hearts || suit == Diamonds {
		return color.Red
	}
	return color.Cyan
}
