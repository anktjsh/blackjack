package main

import (
	"bufio"
	"fmt"
	"github.com/anktjsh/blackjack/gameutil"
	"github.com/gookit/color"
	"os"
	"strings"
)

func PrintIntro() {
	myStyle := color.New(color.FgWhite, color.OpBold)
	myStyle.Println("Welcome to Blackjack!")
	gameutil.PrintPlayingCard(*gameutil.GetJack())
	myStyle.Println("Created by Aniket Joshi")
	myStyle.Println("v1.0.0\n")
}

func main() {
	PrintIntro()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nOptions:")
		color.Info.Println("\tEnter 1 to Play!")
		color.Info.Println("\tEnter 2 to see the Rules")
		color.Info.Println("\tEnter quit anything else to Exit\n")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare(text, "1") == 0 {
			game := new(gameutil.Game)
			game.Deck = new(gameutil.Deck)
			game.Money = gameutil.StartingMoney(reader)
			game.Player = new(gameutil.Hand)
			game.Dealer = new(gameutil.Hand)
			game.Reader = reader
			game.Play()
			if game.Money > 0 {
				color.Info.Println("You ended with", game.Money, "Dollars!")
			} else {
				color.Error.Println("You ended with", game.Money, "Dollars!")
			}
		} else if strings.Compare(text, "2") == 0 {
			gameutil.PrintRules()
		} else if strings.EqualFold(text, "quit") {
			break
		} else {
			color.Error.Println("Not a valid option!")
		}
	}
	color.Info.Println("Thanks for playing! Come again soon!")

}
