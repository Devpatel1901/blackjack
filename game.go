package blackjack

import (
	"fmt"
	"math/rand"

	"slices"

	"github.com/Devpatel1901/cards/v2"
)

func calculateScore(deckOfCards []cards.Card) int {
	total := 0
	numberOfAce := 0
	for _, c := range deckOfCards {
		if c.Rank >= cards.Ten && c.Rank <= cards.King {
			total += 10
		} else if c.Rank == cards.Ace {
			numberOfAce += 1
		} else {
			total = total + int(c.Rank)
		}
	}

	for range numberOfAce {
		if 21-total >= 11 {
			total += 11
		} else {
			total += 1
		}
	}

	return total
}

func initializePlayers() []Player {
	var players []Player

	players = append(players, Player{name: "PLAYER1", isDealer: false}, Player{name: "DEALER", isDealer: true})

	return players
}

func drawCardFromDeck(deckOfCards *[]cards.Card) cards.Card {
	if len(*deckOfCards) == 0 {
		panic("deck is empty")
	}

	idx := rand.Intn(len(*deckOfCards))
	card := (*deckOfCards)[idx]

	*deckOfCards = slices.Delete((*deckOfCards), idx, idx+1)

	return card
}

func dealInitialCards(players []Player, deckOfCards []cards.Card, cardsPerPlayer int) {
	for range cardsPerPlayer {
		for i := range players {
			card := drawCardFromDeck(&deckOfCards)
			players[i].AddCard(card)
		}
	}
}

func showInitialCards(players []Player) {
	for i := range players {
		fmt.Printf("--------------%v CARDS--------------\n", players[i].Name())
		players[i].Show()
	}
}

func hideOneCardFromDealerHands(players []Player) {
	for i := range players {
		if players[i].IsDealer() {
			hands := players[i].Hand()
			idx := rand.Intn(len(hands))

			hands[idx].Hidden = true
		}
	}
}

func StartGame() {
	finalDeck := cards.FromDecks(cards.NewDeck(), cards.NewDeck())

	players := initializePlayers()

	dealInitialCards(players, finalDeck, 2)
	hideOneCardFromDealerHands(players)

	showInitialCards(players)
}
