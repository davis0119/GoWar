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
		if input[0] == 'y' {
			return true
		}
		if input[0] == 'n' {
			return false
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
		d := deck.DeckInit()
		d = deck.Shuffle(d)
		player, bot := deck.SplitCards(d)
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
		gameInSession = promptPlayer("Do you wish to re-wage War? ")
	}
	/*
		game := war.NewGame()
		game.InitPlayer(user)
		game.InitBot(difficulty)
		game.SplitCards()

		for !game.Over() {
			game.CommenceRound() // if card ranks are the same
			 within CommenceRound(), we check for ties and grant the option to forfeit
				for both players like below
				'''
				if isWar {
					playerForfeit := grantForfeitOption()
				}
				'''
			game.PrintBattleResults()
		}
		game.AnnounceWinner()
	*/
}
