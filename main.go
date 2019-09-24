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
	money := int64(0)
initial:
	for {
		fmt.Println("\nOptions:")
		color.Info.Println("\tType 1 to Play!")
		color.Info.Println("\tType 2 to see the Rules")
		color.Info.Println("\tType anything else to Exit\n")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare(text, "1") == 0 {
			deck := new(gameutil.Deck)
			money = gameutil.StartingMoney(reader)

			player := new(gameutil.Hand)
			dealer := new(gameutil.Hand)
			round := 1
			for {
				if money <= 0 {
					color.Error.Println("You ran out of money! The game is over :(")
					break initial
				}
				if money >= 2147483647 {
					color.Info.Println("You have exceeded the limits of this casino! Congratulations!")
					break initial
				}
				deck.Create()

				color.Info.Println("You have", money, "dollars!")
				fmt.Println("Press Enter to continue, Type anything else and then Press Enter to exit")
				str, _ := reader.ReadString('\n')
				if strings.Compare(str, "\n") == 0 {
					color.Info.Println("Round :", round)
					round++
					bet := gameutil.StartingBet(reader, money)
					userEleven := true
					dealerEleven := true
					player.Clear()
					dealer.Clear()
					player.AddCard(deck.PopCard())
					player.AddCard(deck.PopCard())
					dealer.AddCard(deck.PopCard())
					dealer.AddCard(deck.PopCard())
					color.Info.Println("Your hand:")
					player.PrintHand()
					color.Info.Println("Dealer's hand:")
					gameutil.PrintPlayingCard(dealer.GetCard(0))
					color.Info.Println("Dealer's second card is hidden")
					if dealer.IsBlackjack(true) || player.IsBlackjack(true) {
						if dealer.IsBlackjack(true) && player.IsBlackjack(true) {
							dealer.PrintHand()
							color.Warn.Println("Both you and the dealer have Blackjack! Since the hands are tied, no one wins")
						} else if dealer.IsBlackjack(true) {
							color.Error.Println("Dealer has Blackjack! I'm sorry you lose :(")
							dealer.PrintHand()
							money -= bet
						} else {
							color.Info.Println("You have Blackjack! Congratulations, you won :)")
							money += bet
						}
					} else {
						move := true
						for {
							fmt.Println("\nOptions:")
							color.Info.Println("\tType 1 to Hit")
							color.Info.Println("\tType anything else to Stand")
							cardInput, _ := reader.ReadString('\n')
							cardInput = strings.Replace(cardInput, "\n", "", -1)
							if strings.Compare(cardInput, "1") == 0 {
								color.Info.Println("You have decided to get another card")
								player.AddCard(deck.PopCard())
								color.Info.Println("Your hand:")
								player.PrintHand()
								color.Info.Println("Dealer's hand:")
								gameutil.PrintPlayingCard(dealer.GetCard(0))
								color.Info.Println("Dealer's second card is hidden")
								if userEleven {
									if player.IsBust(userEleven) {
										userEleven = false
									}
								}
								if player.IsBust(userEleven) {
									color.Error.Println("Your hand has gone bust! I'm sorry you lose :(")
									money -= bet
									move = false
									break
								}
							} else {
								fmt.Println("You have decided to not get another card")
								break
							}
						}
						if move {
							move = true
							fmt.Println("Dealer's Hand:")
							dealer.PrintHand()
							for dealer.Value(dealerEleven) < 17 {
								color.Warn.Println("Dealer drew a card")
								dealer.AddCard(deck.PopCard())
								fmt.Println("Dealer's Hand:")
								dealer.PrintHand()
								if dealerEleven {
									if dealer.IsBust(dealerEleven) {
										dealerEleven = false
									}
								}
								if dealer.IsBust(userEleven) {
									color.Info.Println("The dealer has gone bust! Congratulations, you won :)")
									money += bet
									move = false
									break
								}
							}
							if move {
								color.Info.Println("Your hand:")
								player.PrintHand()
								if dealer.Value(dealerEleven) > player.Value(userEleven) {
									color.Error.Println("The dealer's hand has a higher value than you! I'm sorry, you lose :(")
									money -= bet
								} else if dealer.Value(dealerEleven) < player.Value(userEleven) {
									color.Info.Println("Your hand is higher than the dealer's! Congratulations, you won :)")
									money += bet
								} else {
									color.Warn.Println("Your hand's are tied, no one wins")
								}
							}
						}
					}
				} else {
					break initial
				}
			}
		} else if strings.Compare(text, "2") == 0 {
			gameutil.PrintRules(reader)
		} else {
			break
		}
	}
	if money > 0 {
		color.Info.Println("You ended with", money, "Dollars!")
	} else {
		color.Error.Println("You ended with", money, "Dollars!")
	}
	color.Info.Println("Thanks for playing! Come again soon!")

}
