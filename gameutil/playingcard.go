package gameutil

import (
	"fmt"
	"github.com/gookit/color"
	"time"
)

var cards = [13]string{"1", "2", "3", "4",
	"5", "6", "7", "8",
	"9", "10", "J", "Q", "K"}

const (
	TEN  Card = 10
	JACK Card = 11
)

type Card int

type PlayingCard struct {
	suit Suit
	card Card
}

func PrintPlayingCard(card PlayingCard) {
	time.Sleep(1 * time.Second)
	printCard(card.card, card.suit)
}

func GetJack() *PlayingCard {
	card := new(PlayingCard)
	card.suit = Hearts
	card.card = JACK
	return card
}

func PrintCards(playing []PlayingCard) {
	for x := 0; x < 9; x++ {
		for i, card := range playing {
			selColor := card.suit.Color()
			myStyle := color.New(selColor, color.OpBold)
			cardStr := cards[card.card-1]
			space := " "
			if card.card == TEN {
				space = ""
			}
			if i > 0 {
				myStyle.Print("\t")
			}
			switch x {
			case 0:
				myStyle.Print("┌─────────┐")
			case 1:
				myStyle.Print(fmt.Sprintf("│%s%s       │", cardStr, space))
			case 2:
				myStyle.Print("│         │")
			case 3:
				myStyle.Print("│         │")
			case 4:
				myStyle.Print(fmt.Sprintf("│    %s    │", card.suit.Visual()))
			case 5:
				myStyle.Print("│         │")
			case 6:
				myStyle.Print("│         │")
			case 7:
				myStyle.Print(fmt.Sprintf("│       %s%s│", space, cardStr))
			case 8:
				myStyle.Print("└─────────┘")
			}
		}
		fmt.Println()
	}
}

func printCard(card Card, suit Suit) {
	var cardList [9]string
	cardStr := cards[card-1]
	space := " "
	if card == TEN {
		space = ""
	}

	cardList[0] = "┌─────────┐"
	cardList[1] = fmt.Sprintf("│%s%s       │", cardStr, space)
	cardList[2] = "│         │"
	cardList[3] = "│         │"
	cardList[4] = fmt.Sprintf("│    %s    │", suit.Visual())
	cardList[5] = "│         │"
	cardList[6] = "│         │"
	cardList[7] = fmt.Sprintf("│       %s%s│", space, cardStr)
	cardList[8] = "└─────────┘"

	selColor := suit.Color()
	myStyle := color.New(selColor, color.OpBold)
	for _, val := range cardList {
		myStyle.Println(val)
	}
}

func (p *PlayingCard) Value(eleven bool) int {
	if p.card >= 10 {
		return 10
	}
	if p.card == 1 {
		if eleven {
			return 11
		}
		return 1
	}
	return int(p.card)

}
