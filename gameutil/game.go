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
	round int
	bet int64
	userEleven bool
	dealerEleven bool
}

func (g* Game) Init(reader *bufio.Reader) {
	g.Deck = new(Deck)
	g.Money = StartingMoney(reader)
	g.Player = new(Hand)
	g.Dealer = new(Hand)
	g.Reader = reader
	g.round = 1
}

func (g* Game) initRound() {
	color.Info.Println("Round :", g.round)
	g.round++
	g.bet = StartingBet(g.Reader, g.Money)
	g.userEleven = true
	g.dealerEleven = true

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
}

func (g* Game) blackjackCheck() bool {
	if g.Dealer.IsBlackjack(true) || g.Player.IsBlackjack(true) {
		if g.Dealer.IsBlackjack(true) && g.Player.IsBlackjack(true) {
			g.Dealer.PrintHand(true)
			color.Warn.Println("Both you and the Dealer have Blackjack! Since the hands are tied, no one wins")
		} else if g.Dealer.IsBlackjack(true) {
			color.Error.Println("Dealer has Blackjack! I'm sorry you lose :(")
			g.Dealer.PrintHand(true)
			g.Money -= g.bet
		} else {
			color.Info.Println("You have Blackjack! Congratulations, you won :)")
			g.Money += g.bet
		}
		return true
	}
	return false
}

func (g* Game) userTurn() bool {
	for {
		fmt.Println("\nOptions:")
		color.Info.Println("\tEnter 1 to Hit")
		color.Info.Println("\tEnter anything else to Stand")
		cardInput, _ := g.Reader.ReadString('\n')
		cardInput = strings.Replace(cardInput, "\n", "", -1)
		if strings.Compare(cardInput, "1") == 0 {
			color.Info.Println("You have decided to get another card")
			g.Player.AddCard(g.Deck.PopCard())

			if g.userEleven {
				if g.Player.IsBust(g.userEleven) {
					g.userEleven = false
				}
			}

			color.Info.Println("Your hand:")
			g.Player.PrintHand(g.userEleven)
			color.Info.Println("Dealer's hand:")
			PrintPlayingCard(g.Dealer.GetCard(0))
			color.Info.Println("Dealer's second card is hidden")

			if g.Player.IsBust(g.userEleven) {
				color.Error.Println("Your hand has gone bust! I'm sorry you lose :(")
				g.Money -= g.bet
				return true
			}
		} else {
			color.Info.Println("You have decided not to get another card")
			break
		}
	}
	return false
}

func (g* Game) dealerTurn() bool {
	color.Info.Println("Dealer's Hand:")
	g.Dealer.PrintHand(g.dealerEleven)
	for g.Dealer.Value(g.dealerEleven) < 17 {
		color.Warn.Println("Dealer drew a card")
		g.Dealer.AddCard(g.Deck.PopCard())

		if g.dealerEleven {
			if g.Dealer.IsBust(g.dealerEleven) {
				g.dealerEleven = false
			}
		}

		color.Info.Println("Dealer's Hand:")
		g.Dealer.PrintHand(g.dealerEleven)

		if g.Dealer.IsBust(g.dealerEleven) {
			color.Info.Println("The Dealer has gone bust! Congratulations, you won :)")
			g.Money += g.bet
			return true
		}
	}
	return false
}

func (g* Game) finalCheck() {
	color.Info.Println("Your hand:")
	g.Player.PrintHand(g.userEleven)
	if g.Dealer.Value(g.dealerEleven) > g.Player.Value(g.userEleven) {
		color.Error.Println("The Dealer's hand has a higher value than you! I'm sorry, you lose :(")
		g.Money -= g.bet
	} else if g.Dealer.Value(g.dealerEleven) < g.Player.Value(g.userEleven) {
		color.Info.Println("Your hand is higher than the Dealer's! Congratulations, you won :)")
		g.Money += g.bet
	} else {
		color.Warn.Println("Your hand's are tied, no one wins")
	}
}

func (g* Game) Play() {
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
			g.initRound()
			if !g.blackjackCheck() {
				if !g.userTurn() {
					if !g.dealerTurn() {
						g.finalCheck()
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
