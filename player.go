package blackjack

import (
	"fmt"

	"github.com/Devpatel1901/cards/v2"
)

type Player struct {
	name     string
	isDealer bool
	hand     []cards.Card
}

func (p Player) Name() string {
	return p.name
}

func (p Player) IsDealer() bool {
	return p.isDealer
}

func (p Player) Show() {
	for _, c := range p.hand {
		fmt.Println(c.Print())
	}
}

func (p Player) Hand() []cards.Card {
	return p.hand
}

func (p *Player) AddCard(cardsToAdd ...cards.Card) {
	p.hand = append(p.hand, cardsToAdd...)
}
