package main

import (
	"fmt"

	"./deck"
)

// Input: string
// Return: bool
// Prints the given prompt and scans for an input. Continually checks
// the first character of the input and returns true if the input is 'y' or
// false if the input is 'n'.
func promptPlayer(prompt string) bool {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scanln(&input)
		if len(input) > 0 {
			if input[0] == 'y' {
				return true
			}
			if input[0] == 'n' {
				return false
			}
			if input == "help" {
				deck.FocusTerminal()
				fmt.Println("War is a card game which involves 2 players. Each player receives half of a shuffled deck of cards.")
				fmt.Println("Each round, you will be prompted to commence a battle. In a battle, each player draws a card. The player with a higher card rank (ignoring suit) wins the battle and adds all cards played to their deck. If the ranks are the same, a War Event occurs. You may reprompt 'help' then for more info about that case.")
				fmt.Println("A winner is determined when one side no longer has any more cards left to play.")
			}
		}
		if input != "help" {
			fmt.Println("Answer via 'y' or 'n'. If you would like more info, type 'help'")
		}
	}
}

// Input: deck pointer, deck pointer
// Prints out the current status of the game (how many cards each player has).
func announceStatusQuo(player *deck.Deck, bot *deck.Deck) {
	fmt.Println("You:", len(player.PlayingCards), "cards |", "Opponent:", len(bot.PlayingCards), "cards")
}

// Input: deck pointer, deck pointer
// Announces a winner if there is one.
func announceWinner(player *deck.Deck, bot *deck.Deck) {
	if len(player.PlayingCards) == 0 {
		fmt.Println("YOU HAVE LOST THE WAR!")
	} else if len(bot.PlayingCards) == 0 {
		fmt.Println("YOU HAVE WON THE WAR!")
	}
}

func main() {
	gameInSession := promptPlayer("Do you wish to enter the War? (y/n | 'help' for more info) ")
	for gameInSession {
		deck.FocusTerminal()
		// Prepare the deck of cards.
		d := deck.DeckInit()
		d = deck.Shuffle(d)
		// Hand out the cards.
		player, bot := deck.SplitCards(d)
		// Game Loop.
		for !deck.GameOver(player, bot) {
			declareBattle := promptPlayer("Wage battle? (y/n) ")
			if declareBattle {
				deck.CommenceRound(player, bot)
				announceStatusQuo(player, bot)
			} else {
				break
			}
		}
		fmt.Println("The game is over.")
		announceStatusQuo(player, bot)
		announceWinner(player, bot)
		// Restart game option.
		gameInSession = promptPlayer("Do you wish to re-wage War? ")
	}
}
