package gameutil

import "time"

type Hand struct {
	cards []PlayingCard
}

func (h *Hand) AddCard(p PlayingCard) {
	h.cards = append(h.cards, p)
}

func (h *Hand) IsBlackjack(eleven bool) bool {
	if len(h.cards) == 2 {
		if h.cards[0].Value(eleven)+h.cards[1].Value(eleven) == 21 {
			return true
		}
	}
	return false
}

func (h *Hand) IsBust(eleven bool) bool {
	return h.Value(eleven) > 21
}

func (h *Hand) Value(eleven bool) int {
	sum := 0
	for i := 0; i < len(h.cards); i++ {
		sum += int(h.cards[i].Value(eleven))
	}
	return sum
}

func (h *Hand) Size() int {
	return len(h.cards)
}

func (h *Hand) GetCard(index int) PlayingCard {
	return h.cards[index]
}

func (h *Hand) Clear() {
	h.cards = h.cards[:0]
}

func (h *Hand) PrintHand() {
	time.Sleep(1 * time.Second)
	PrintCards(h.cards)
}
