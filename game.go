package blackjack

func StartGame(numberOfPlayers int, minBet, maxBet int) {

	gs := CreateNewGame(numberOfPlayers, minBet, maxBet)

	checkForNaturalBlackJack(gs)

	playPlayerTurn(gs)

	playDealerTurn(gs)

	endGameAndDisplayStats(gs)
}
