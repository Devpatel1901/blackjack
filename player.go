package blackjack

import (
	"fmt"

	"github.com/Devpatel1901/cards/v2"
)

type Player struct {
	name string
	hand []cards.Card
	bet  int
}

func (p Player) Name() string {
	return p.name
}

func (p Player) Show() {
	for _, c := range p.hand {
		fmt.Println(c.Print())
	}
}

func (p Player) Hand() []cards.Card {
	return p.hand
}

func (p Player) Score() (int, bool) {
	total := 0
	numberOfAce := 0
	softScore := false

	for _, c := range p.hand {
		switch {
		case c.Rank >= cards.Ten && c.Rank <= cards.King:
			total += 10
		case c.Rank == cards.Ace:
			numberOfAce++
			total += 1
		default:
			total += int(c.Rank)
		}
	}

	if numberOfAce > 0 && total+10 <= 21 {
		total += 10
		softScore = true
	}

	return total, softScore
}

func (p *Player) AddCard(cardsToAdd ...cards.Card) {
	p.hand = append(p.hand, cardsToAdd...)
}

func (p *Player) IncreaseBetByAmount(amount int) {
	p.bet += amount
}
