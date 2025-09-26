# CLI Blackjack Game

Welcome to the CLI Blackjack Game! This project is a command-line implementation of the classic Blackjack card game. Follow the instructions below to set up and play the game.

---

## Project Setup

### Prerequisites
- Go 1.20 or higher installed on your system.
- Basic knowledge of how to use the terminal.

### Steps to Set Up
1. Clone the repository:
   ```bash
   git clone https://github.com/Devpatel1901/blackjack.git
   ```
2. Navigate to the project directory:
   ```bash
   cd blackjack
   ```
3. Build the project:
   ```bash
   go build -o main/main main/main.go
   ```
4. Ensure the binary file is created at `main/main`.

---

## How to Start the Game

1. Open your terminal and navigate to the project directory.
2. Run the help command to know more about command line arguments:
   ```bash
   ./main/main --help
   ```
   Output:
   ```code
   Usage of ./main/main:
        -maxBet int
            Maximum betting amount required to play on this table. (default 1000)
        -minBet int
            Minimum betting amount required to play on this table. (default 10)
        -players int
            Number of Players (except dealer) on one table. (default 2)
   ```
2. Run the binary file to start the game:
   ```bash
   ./main/main -players=2 -minBet=10 -maxBet=1000
   ```
3. Follow the on-screen instructions to play the game.

---

## Game Overview

### Objective
The goal of the game is to get as close to 21 as possible without exceeding it. The dealer must hit on 16 and stand on 17. A natural Blackjack (an Ace and a 10-value card) pays immediately.

### Starting the Game
1. **Welcome Screen**: The game begins with a welcome banner and a brief overview of the rules.
2. **Player Setup**: Enter the number of players and place your bets. Each player must bet an amount within the table's minimum and maximum limits.
3. **Initial Deal**: Each player and the dealer are dealt two cards. One of the dealer's cards remains hidden.

### Player Turn
- Each player takes turns to play.
- On your turn, you can:
  - **Hit**: Draw another card.
  - **Stand**: Keep your current hand.
  - **Double Down**: Double your bet and draw one final card.

### Dealer Turn
- The dealer reveals their hidden card and plays according to the rules:
  - Hit on 16 or less.
  - Stand on 17 or higher.

### End of Game
- The game ends when all players and the dealer have completed their turns.
- The winner is determined based on the scores:
  - Closest to 21 without exceeding it.
  - If all players bust, the dealer wins.

---

## Example Game Iteration

1. **Start the Game**:
   ```bash
   ./main/main
   ```
2. **Place Bets**:
   - Each player places their bets within the table limits.
3. **Initial Deal**:
   - Players and the dealer receive their initial cards.
4. **Player Turns**:
   - Each player decides to hit, stand, or double down.
5. **Dealer Turn**:
   - The dealer plays according to the rules.
6. **Results**:
   - The winner is announced, and the game displays the final scores and bets.

---

## Some Snippets of Blackjack Game

![Intro Image](<images/intro.png>)

![Turn Image](<images/turn.png>)

![Result Image](<images/result.png>)

Enjoy playing CLI Blackjack! 🎉

> Developed by [Dev Patel](mailto:entreprenerudev1901@gmail.com)

> If you are interested and want to colloborate or know more about this project connect with me on [LinkedIn](https://www.linkedin.com/in/dev-patel-540a84233/)

> Stay tune for upgrade!!!