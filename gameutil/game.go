package gameutil

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"strings"
)

type Game struct {
	Deck *Deck
	Money int64
	Player *Hand
	Dealer *Hand
	Reader *bufio.Reader
}

func (g* Game) Play() {
	round := 1
	for {
		if g.Money <= 0 {
			color.Error.Println("You ran out of Money! The game is over :(")
			return
		}
		if g.Money >= 2147483647 {
			color.Info.Println("You have exceeded the limits of this casino! Congratulations!")
			return
		}
		g.Deck.Create()

		color.Info.Println("You have", g.Money, "dollars!")
		fmt.Println("Press Enter to continue, Enter quit to exit")
		str, _ := g.Reader.ReadString('\n')
		str = strings.Replace(str, "\n", "", -1)
		if strings.Compare(str, "") == 0 {
			color.Info.Println("Round :", round)
			round++
			bet := StartingBet(g.Reader, g.Money)
			userEleven := true
			dealerEleven := true
			g.Player.Clear()
			g.Dealer.Clear()
			g.Player.AddCard(g.Deck.PopCard())
			g.Player.AddCard(g.Deck.PopCard())
			g.Dealer.AddCard(g.Deck.PopCard())
			g.Dealer.AddCard(g.Deck.PopCard())
			color.Info.Println("Your hand:")
			g.Player.PrintHand(true)
			color.Info.Println("Dealer's hand:")
			PrintPlayingCard(g.Dealer.GetCard(0))
			color.Info.Println("Dealer's second card is hidden")
			if g.Dealer.IsBlackjack(true) || g.Player.IsBlackjack(true) {
				if g.Dealer.IsBlackjack(true) && g.Player.IsBlackjack(true) {
					g.Dealer.PrintHand(true)
					color.Warn.Println("Both you and the Dealer have Blackjack! Since the hands are tied, no one wins")
				} else if g.Dealer.IsBlackjack(true) {
					color.Error.Println("Dealer has Blackjack! I'm sorry you lose :(")
					g.Dealer.PrintHand(true)
					g.Money -= bet
				} else {
					color.Info.Println("You have Blackjack! Congratulations, you won :)")
					g.Money += bet
				}
			} else {
				move := true
				for {
					fmt.Println("\nOptions:")
					color.Info.Println("\tEnter 1 to Hit")
					color.Info.Println("\tEnter anything else to Stand")
					cardInput, _ := g.Reader.ReadString('\n')
					cardInput = strings.Replace(cardInput, "\n", "", -1)
					if strings.Compare(cardInput, "1") == 0 {
						color.Info.Println("You have decided to get another card")
						g.Player.AddCard(g.Deck.PopCard())

						if userEleven {
							if g.Player.IsBust(userEleven) {
								userEleven = false
							}
						}

						color.Info.Println("Your hand:")
						g.Player.PrintHand(userEleven)
						color.Info.Println("Dealer's hand:")
						PrintPlayingCard(g.Dealer.GetCard(0))
						color.Info.Println("Dealer's second card is hidden")

						if g.Player.IsBust(userEleven) {
							color.Error.Println("Your hand has gone bust! I'm sorry you lose :(")
							g.Money -= bet
							move = false
							break
						}
					} else {
						fmt.Println("You have decided not to get another card")
						break
					}
				}
				if move {
					move = true
					fmt.Println("Dealer's Hand:")
					g.Dealer.PrintHand(dealerEleven)
					for g.Dealer.Value(dealerEleven) < 17 {
						color.Warn.Println("Dealer drew a card")
						g.Dealer.AddCard(g.Deck.PopCard())

						if dealerEleven {
							if g.Dealer.IsBust(dealerEleven) {
								dealerEleven = false
							}
						}

						fmt.Println("Dealer's Hand:")
						g.Dealer.PrintHand(dealerEleven)

						if g.Dealer.IsBust(userEleven) {
							color.Info.Println("The Dealer has gone bust! Congratulations, you won :)")
							g.Money += bet
							move = false
							break
						}
					}
					if move {
						color.Info.Println("Your hand:")
						g.Player.PrintHand(userEleven)
						if g.Dealer.Value(dealerEleven) > g.Player.Value(userEleven) {
							color.Error.Println("The Dealer's hand has a higher value than you! I'm sorry, you lose :(")
							g.Money -= bet
						} else if g.Dealer.Value(dealerEleven) < g.Player.Value(userEleven) {
							color.Info.Println("Your hand is higher than the Dealer's! Congratulations, you won :)")
							g.Money += bet
						} else {
							color.Warn.Println("Your hand's are tied, no one wins")
						}
					}
				}
			}
		} else if strings.EqualFold(str, "quit") {
			return
		} else {
			color.Error.Println("Not a valid option!")
		}
	}
}
