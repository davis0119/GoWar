package main

import (
	"fmt"

	"./deck"
)

func enterGame(prompt string) bool {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scanln(&input)
		if input[0] == 'y' {
			return true
		} else if input[0] == 'n' {
			return false
		}
		fmt.Println("Answer via 'y' or 'n'")
	}
}

func main() {
	// initialize deck variable
	d := deck.DeckInit()
	gameInSession := enterGame("Do you wish to enter War? ")
	for gameInSession {
		// deck.SplitCards(d)
		d1, d2 := deck.SplitCards(d)
		// print("Length of the deck: ", len(d.PlayingCards))
		println("Length of d1: ", len(d1.PlayingCards))
		println("Length of d2: ", len(d2.PlayingCards))
		gameInSession = false
	}
	// deal out cards to both players

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
