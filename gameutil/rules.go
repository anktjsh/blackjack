package gameutil

import (
	"github.com/gookit/color"
	"strconv"
	"time"
)

var rules = []string{
	"The goal of blackjack is to beat the dealer's hand without going over 21",
	"Face cards are worth 10. Aces are worth 1 or 11, whichever makes a better hand",
	"Each player starts with two cards, one of the dealer's cards is hidden until the end",
	"To 'Hit' is to ask for another card. To 'Stand' is to hold your total and end your turn",
	"If you go over 21 you bust, and the dealer wins regardless of the dealer's hand",
	"Dealer will hit until his/her cards total 17 or higher",
	"To start each round, you bet a certain amount of money",
	"If you win the hand, you will earn the amount of money you bet on the hand",
}

func PrintRules() {
	color.Magenta.Println("\nBasic Blackjack Rules:")
	for i := 0; i < len(rules); i++ {
		time.Sleep(1*time.Second)
		color.Info.Println(strconv.Itoa(i+1)+".", rules[i])
	}
}
