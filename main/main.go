package main

import (
	"flag"

	"github.com/Devpatel1901/blackjack"
)

func main() {
	numberOfPlayers := flag.Int("players", 2, "Number of Players (except dealer) on one table.")
	minBet := flag.Int("minBet", 10, "Minimum betting amount required to play on this table.")
	maxBet := flag.Int("maxBet", 1000, "Maximum betting amount required to play on this table.")

	flag.Parse()

	blackjack.StartGame(*numberOfPlayers, *minBet, *maxBet)
}
