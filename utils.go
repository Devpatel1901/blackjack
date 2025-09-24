package blackjack

import (
	"fmt"

	"github.com/Devpatel1901/cards/v2"
)

func showPlayerCards(players ...Player) {
	for _, p := range players {
		fmt.Printf("-------------- %v CARDS --------------\n", p.Name())
		p.Show()
	}
}

func draw(deck []cards.Card) (cards.Card, []cards.Card, error) {
	if len(deck) == 0 {
		return cards.Card{}, deck, fmt.Errorf("deck is empty")
	}
	return deck[0], deck[1:], nil
}

func initializePlayers(numberOfPlayers int) []Player {
	var players []Player

	for i := range numberOfPlayers {
		players = append(players, Player{name: fmt.Sprintf("PLAYER%d", i+1)})
	}

	return players
}

func dealInitialCards(players []Player, deck []cards.Card) ([]Player, []cards.Card) {
	var card cards.Card
	var err error

	for j := 0; j < 2; j++ {
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

func dealDealerInitialCards(dealer Player, deck []cards.Card) (Player, []cards.Card) {
	var card cards.Card
	var err error

	card, deck, err = draw(deck)
	if err == nil {
		card.Hidden = true
		dealer.AddCard(card)
	}

	card, deck, err = draw(deck)
	if err == nil {
		card.Hidden = false
		dealer.AddCard(card)
	}

	return dealer, deck
}

func revealDealerFirstHiddenCard(gs *GameState) {
	for i := range gs.Dealer.hand {
		if gs.Dealer.hand[i].Hidden {
			gs.Dealer.hand[i].Hidden = false
			break
		}
	}
}

func placeBets(players []Player, minBet int, maxBet int) int {
	totalBetAmount := 0
	for i := range players {
		input := 0

		for {
			fmt.Printf("Table minimum allowed bet is: $%d and maximum allowed bet is: $%d. Enter your bet %v: $", minBet, maxBet, players[i].Name())
			fmt.Scanf("%d\n", &input)

			if input < minBet || input > maxBet {
				fmt.Println("Try Again, Bet amount is invalid!!!")
				continue
			} else {
				totalBetAmount += input
				players[i].IncreaseBetByAmount(input)
				break
			}
		}
	}
	return totalBetAmount
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

func playPlayerTurn(gs GameState) {
	var input string

	for i := range len(gs.Players) {
		for input != "s" {
			showPlayerCards(gs.Players...)
			showPlayerCards(gs.Dealer)

			fmt.Printf("What will you do %v? (h)it, (s)tand, (d)ouble down: ", gs.Players[i].Name())
			fmt.Scanf("%s\n", &input)

			var card cards.Card
			var err error

			switch input {
			case "h":
				card, gs.Deck, err = draw(gs.Deck)
				if err != nil {
					fmt.Println("No more cards left in deck!")
					break
				}
				gs.Players[i].AddCard(card)
			case "d":
				amount := 0
				for {
					fmt.Printf("Table minimum allowed bet is: $%d and maximum allowed bet is: $%d. Enter your bet %v: $", gs.MinTableBet, gs.MaxTableBet, gs.Players[i].Name())
					fmt.Scanf("%d\n", &amount)

					if amount < gs.MinTableBet || amount > gs.MaxTableBet {
						fmt.Println("Try Again, Bet amount is invalid!!!")
						continue
					} else {
						gs.TotalBet += amount
						gs.Players[i].IncreaseBetByAmount(amount)
						break
					}
				}

				card, gs.Deck, err = draw(gs.Deck)
				if err != nil {
					fmt.Println("No more cards left in deck!")
					break
				}
				gs.Players[i].AddCard(card)
			}
		}
		input = ""
	}
}

func playDealerTurn(gs GameState) {
	var card cards.Card
	var err error

	dealerScore, isSoftScore := gs.Dealer.Score()

	for dealerScore <= 16 || (dealerScore == 17 && isSoftScore) {
		card, gs.Deck, err = draw(gs.Deck)
		if err != nil {
			fmt.Println("No more cards left in deck!")
			break
		}
		gs.Dealer.AddCard(card)

		dealerScore, isSoftScore = gs.Dealer.Score()
	}
}

func endGameAndDisplayStats(gs GameState) {
	revealDealerFirstHiddenCard(&gs)

	showPlayerCards(gs.Players...)
	showPlayerCards(gs.Dealer)
	fmt.Printf("--------------- MATCH RESULT ---------------\n")

	var maxScore int
	var winner string

	for i := range gs.Players {
		pScore, _ := gs.Players[i].Score()

		if pScore > 21 {
			fmt.Printf("%v busted\n", gs.Players[i].Name())
			continue
		}

		if pScore > maxScore {
			maxScore = pScore
			winner = gs.Players[i].Name()
		}
	}

	fmt.Printf("%v win!!!\n", winner)
	fmt.Printf("Total Winning Amount: %d\n", gs.TotalBet)
}
