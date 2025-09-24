package blackjack

import (
	"fmt"
	"os"
)

func StartGame() {
	gs := CreateNewGame(4, 10, 1000)

	for i := range len(gs.Players) {
		if hasNaturalBlackjack(gs.Players[i].Hand()) {
			showPlayerCards(gs.Players...)
			showPlayerCards(gs.Dealer)
			fmt.Printf("***************%v HAS A NATURAL BLACKJACK, SO %v IS AN IMMEDIATE WINNER***************\n", gs.Players[i].Name(), gs.Players[i].Name())
			fmt.Printf("Total Winning Amount: %d\n", gs.TotalBet)
			os.Exit(0)
		}
	}

	playPlayerTurn(gs)

	playDealerTurn(gs)

	endGameAndDisplayStats(gs)
}
