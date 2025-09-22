package blackjack

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/Devpatel1901/cards/v2"
)

func initializePlayers(numberOfPlayers int) []Player {
	var players []Player

	for i := range numberOfPlayers {
		players = append(players, Player{name: fmt.Sprintf("PLAYER%d", i+1), isDealer: false})
	}

	players = append(players, Player{name: "DEALER", isDealer: true})

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

func setDealerCardVisibility(players []Player, hidden bool) {
	for i := range players {
		if players[i].IsDealer() {
			hands := players[i].Hand()

			if hidden {
				idx := rand.Intn(len(hands))
				hands[idx].Hidden = true
				return
			}

			for j := range hands {
				if hands[j].Hidden {
					hands[j].Hidden = false
					return
				}
			}
		}
	}
}

func dealInitialCards(players []Player, deck []cards.Card, cardsPerPlayer int) ([]Player, []cards.Card) {
	var card cards.Card
	var err error

	for j := 0; j < cardsPerPlayer; j++ {
		for i := range players {
			card, deck, err = draw(deck)
			if err != nil {
				fmt.Println("No more cards left in deck!")
				return players, deck
			}
			players[i].AddCard(card)
		}
	}
	return players, deck
}

func hasNaturalBlackjack(hand []cards.Card) bool {
	if len(hand) != 2 {
		return false
	}

	first := hand[0].Rank.Single()
	second := hand[1].Rank.Single()

	isAceFirst := first == "A" && (second == "10" || second == "J" || second == "Q" || second == "K")
	isAceSecond := second == "A" && (first == "10" || first == "J" || first == "Q" || first == "K")

	return isAceFirst || isAceSecond
}

func StartGame() {
	deckOfCards := cards.FromDecks(cards.NewDeck(cards.Shuffle), cards.NewDeck(cards.Shuffle))
	players := initializePlayers(4)

	var card cards.Card
	var err error
	players, deckOfCards = dealInitialCards(players, deckOfCards, 2)

	for i := range len(players) {
		if hasNaturalBlackjack(players[i].Hand()) {
			fmt.Printf("%v win!!!", players[i].Name())
			os.Exit(0)
		}
	}

	setDealerCardVisibility(players, true)

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

		dealerScore, isSoftScore = players[len(players)-1].Score()
	}

	setDealerCardVisibility(players, true)

	showPlayerCards(players)

	fmt.Printf("--------------- MATCH RESULT ---------------\n")

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
