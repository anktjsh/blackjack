package gameutil

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"strconv"
	"strings"
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

func PrintRules(reader *bufio.Reader) {
	//enter to continue, q to quit
	//improve this
	color.Magenta.Println("\nBasic Blackjack Rules:")
	for i := 0; i < len(rules); i++ {
		color.Info.Println(strconv.Itoa(i+1)+".", rules[i])
		fmt.Println("Press Enter to continue, Type anything else and then Press Enter to exit")
		str, _ := reader.ReadString('\n')
		if strings.Compare(str, "\n") == 0 {
			continue
		} else {
			break
		}
	}
}
