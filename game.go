package blackjack

import (
	"fmt"
	"math/rand"

	"github.com/Devpatel1901/cards/v2"
)

func initializePlayers() []Player {
	var players []Player

	players = append(players, Player{name: "PLAYER1", isDealer: false}, Player{name: "DEALER", isDealer: true})

	return players
}

func draw(deck []cards.Card) (cards.Card, []cards.Card, error) {
	if len(deck) == 0 {
		return cards.Card{}, deck, fmt.Errorf("deck is empty")
	}
	return deck[0], deck[1:], nil
}

func showPlayerCards(players []Player) {
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
	deckOfCards := cards.FromDecks(cards.NewDeck(cards.Shuffle), cards.NewDeck(cards.Shuffle))
	fmt.Println(len(deckOfCards))
	players := initializePlayers()

	var card cards.Card
	var err error
	for j := 0; j < 2; j++ {
		for i := range players {
			card, deckOfCards, err = draw(deckOfCards)
			if err != nil {
				fmt.Println("No more cards left in deck!")
				break
			}
			players[i].AddCard(card)
		}
	}

	hideOneCardFromDealerHands(players)

	var input string

	for i := range len(players) - 1 {
		for input != "s" {
			showPlayerCards(players)

			fmt.Printf("What will you do %v? (h)it, (s)tand: ", players[i].Name())
			fmt.Scanf("%s\n", &input)

			switch input {
			case "h":
				card, deckOfCards, err = draw(deckOfCards)
				if err != nil {
					fmt.Println("No more cards left in deck!")
					break
				}
				players[i].AddCard(card)
			}
		}
		input = ""
	}

	dealerScore, isSoftScore := players[len(players)-1].Score()

	for dealerScore <= 16 || (dealerScore == 17 && isSoftScore) {
		card, deckOfCards, err = draw(deckOfCards)
		if err != nil {
			fmt.Println("No more cards left in deck!")
			break
		}
		players[len(players)-1].AddCard(card)
	}

	var maxScore int
	var winner string

	for i := range players {
		pScore, _ := players[i].Score()

		if pScore > 21 {
			fmt.Printf("%v busted\n", players[i].Name())
			continue
		}

		if pScore > maxScore {
			maxScore = pScore
			winner = players[i].Name()
		}
	}

	fmt.Printf("%v win!!!", winner)
}
