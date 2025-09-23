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
