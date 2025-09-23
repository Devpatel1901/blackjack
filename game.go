package blackjack

import (
	"fmt"
	"os"
)

func StartGame() {
	gs := CreateNewGame(4, 10, 1000)

	showPlayerCards(gs.Players...)
	showPlayerCards(gs.Dealer)

	for i := range len(gs.Players) {
		if hasNaturalBlackjack(gs.Players[i].Hand()) {
			fmt.Printf("***************%v HAS A NATURAL BLACKJACK, SO %v IS AN IMMEDIATE WINNER***************\n", gs.Players[i].Name(), gs.Players[i].Name())
			fmt.Printf("Total Winning Amount: %d\n", gs.TotalBet)
			os.Exit(0)
		}
	}

	// var input string

	// for i := range len(players) - 1 {
	// 	for input != "s" {
	// 		showPlayerCards(players)

	// 		fmt.Printf("What will you do %v? (h)it, (s)tand, (d)ouble down: ", players[i].Name())
	// 		fmt.Scanf("%s\n", &input)

	// 		switch input {
	// 		case "h":
	// 			card, deckOfCards, err = draw(deckOfCards)
	// 			if err != nil {
	// 				fmt.Println("No more cards left in deck!")
	// 				break
	// 			}
	// 			players[i].AddCard(card)
	// 		case "d":
	// 			amount := 0
	// 			for {
	// 				fmt.Printf("Table minimum allowed bet is: $%d and maximum allowed bet is: $%d. Enter your bet %v: $", minBet, maxBet, players[i].Name())
	// 				fmt.Scanf("%d\n", &amount)

	// 				if amount < minBet || amount > maxBet {
	// 					fmt.Println("Try Again, Bet amount is invalid!!!")
	// 					continue
	// 				} else {
	// 					totalBetAmount += amount
	// 					players[i].IncreaseBetByAmount(amount)
	// 					break
	// 				}
	// 			}

	// 			card, deckOfCards, err = draw(deckOfCards)
	// 			if err != nil {
	// 				fmt.Println("No more cards left in deck!")
	// 				break
	// 			}
	// 			players[i].AddCard(card)
	// 		}
	// 	}
	// 	input = ""
	// }

	// dealerScore, isSoftScore := players[len(players)-1].Score()

	// for dealerScore <= 16 || (dealerScore == 17 && isSoftScore) {
	// 	card, deckOfCards, err = draw(deckOfCards)
	// 	if err != nil {
	// 		fmt.Println("No more cards left in deck!")
	// 		break
	// 	}
	// 	players[len(players)-1].AddCard(card)

	// 	dealerScore, isSoftScore = players[len(players)-1].Score()
	// }

	// setDealerCardVisibility(players, false)

	// showPlayerCards(players)

	// fmt.Printf("--------------- MATCH RESULT ---------------\n")

	// var maxScore int
	// var winner string

	// for i := range players {
	// 	pScore, _ := players[i].Score()

	// 	if pScore > 21 {
	// 		fmt.Printf("%v busted\n", players[i].Name())
	// 		continue
	// 	}

	// 	if pScore > maxScore {
	// 		maxScore = pScore
	// 		winner = players[i].Name()
	// 	}
	// }

	// fmt.Printf("%v win!!!\n", winner)
	fmt.Printf("Total Winning Amount: %d\n", gs.TotalBet)
}
