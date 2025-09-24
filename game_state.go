package blackjack

import (
	"github.com/Devpatel1901/cards/v2"
)

type GameState struct {
	Deck        []cards.Card
	Players     []Player
	Dealer      Player
	MinTableBet int
	MaxTableBet int
	TotalBet    int
}

func CreateNewGame(numberOfPlayers int, minTableBet int, maxTableBet int) GameState {
	printWelcomeBanner()
	deck := cards.FromDecks(
		cards.NewDeck(cards.Shuffle),
		cards.NewDeck(cards.Shuffle),
	)

	players := initializePlayers(numberOfPlayers)
	dealer := Player{name: "DEALER"}

	totalBet := placeBets(players, minTableBet, maxTableBet)

	players, deck = dealInitialCards(players, deck)
	dealer, deck = dealDealerInitialCards(dealer, deck)

	gs := GameState{
		Deck:        deck,
		Players:     players,
		Dealer:      dealer,
		MinTableBet: minTableBet,
		MaxTableBet: maxTableBet,
		TotalBet:    totalBet,
	}

	return gs
}
