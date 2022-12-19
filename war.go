package main

import (
	"fmt"

	"./deck"
)

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

func gameOver(player *deck.Deck, bot *deck.Deck) bool {
	return len(player.PlayingCards) == 0 || len(bot.PlayingCards) == 0
}

func AnnounceStatusQuo(player *deck.Deck, bot *deck.Deck) {
	fmt.Println("You:", len(player.PlayingCards), "cards |", "Opponent:", len(bot.PlayingCards), "cards")
}

func AnnounceWinner(player *deck.Deck, bot *deck.Deck) {
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
				AnnounceStatusQuo(player, bot)
			} else {
				break
			}
		}
		fmt.Println("The game is over.")
		AnnounceStatusQuo(player, bot)
		AnnounceWinner(player, bot)
		// Restart game option.
		gameInSession = promptPlayer("Do you wish to re-wage War? ")
	}
}
