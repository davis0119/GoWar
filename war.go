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
		}
		fmt.Println("Answer via 'y' or 'n'")
	}
}

// Input: deck pointer, deck pointer
// Checks if either of the decks have 0 cards and returns true if either one does.
func gameOver(player *deck.Deck, bot *deck.Deck) bool {
	return len(player.PlayingCards) == 0 || len(bot.PlayingCards) == 0
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
	gameInSession := promptPlayer("Do you wish to enter War? ")
	for gameInSession {
		for i := 0; i < 10; i++ {
			println() // do this to reduce clutter on the screen
		}
		// Prepare the deck of cards.
		d := deck.DeckInit()
		d = deck.Shuffle(d)
		// Hand out the cards.
		player, bot := deck.SplitCards(d)
		// Game Loop.
		for !gameOver(player, bot) {
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
