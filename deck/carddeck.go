package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit struct {
	Name string
}

type Value struct {
	Name string
	Val  int
}

type Card struct {
	Suit  Suit
	Value Value
}

type Deck struct {
	PlayingCards []Card
}

func (c *Card) ToStr() string {
	return c.Value.Name + " of " + c.Suit.Name
}

func Shuffle(d *Deck) *Deck {
	rand.Seed(time.Now().UnixNano()) // not doing this doesn't really "shuffle" the deck every time
	for i := len(d.PlayingCards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.PlayingCards[i], d.PlayingCards[j] = d.PlayingCards[j], d.PlayingCards[i] // swap the cards
	}
	return d
}

func Draw(d *Deck) (Card, *Deck) {
	card := d.PlayingCards[0]
	d.PlayingCards = d.PlayingCards[1:]
	return card, d
}

func SplitCards(d *Deck) (d1 *Deck, d2 *Deck) {
	d1 = new(Deck)
	d2 = new(Deck)
	for len(d.PlayingCards) > 0 {
		c1, _ := Draw(d)
		c2, _ := Draw(d)
		d1.PlayingCards = append(d1.PlayingCards, c1)
		d2.PlayingCards = append(d2.PlayingCards, c2)
	}
	return d1, d2
}

func (d *Deck) ToStr() string {
	s := ""
	for _, c := range d.PlayingCards {
		println(c.ToStr())
	}
	return s
}

func DeckInit() *Deck {
	d := new(Deck)
	suits := []Suit{
		{"Spades"},
		{"Clubs"},
		{"Diamonds"},
		{"Hearts"},
	}
	values := []Value{
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"10", 10},
		{"Jack", 11},
		{"Queen", 12},
		{"King", 13},
		{"Ace", 14},
	}
	for _, v := range values {
		for _, s := range suits {
			d.PlayingCards = append(d.PlayingCards, Card{Suit: s, Value: v})
		}
	}
	return d
}

func (c1 *Card) BattleAgainst(c2 Card) int {
	fmt.Println("You: " + c1.ToStr() + " | Opponent: " + c2.ToStr())
	if c1.Value.Val < c2.Value.Val {
		return -1 // loses
	} else if c1.Value.Val == c2.Value.Val {
		return 0 // tie (WAR)
	} else {
		return 1 // win
	}
}

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
				fmt.Println("War Event: Occurs due to both players drawing the same rank card.")
				fmt.Println("In this event, both players place down 3 cards and draw one additional card to battle.")
				fmt.Println("The winner takes all cards on the field.")
				fmt.Println("If there is another tie, this War Event recursively continues. If a player runs out of cards, they lose the game.")
			}
		}
		if input != "help" {
			fmt.Println("Answer via 'y' or 'n'")
		}
	}
}

// The event when 2 cards are played with equal rank.
// Each player draws 3 cards from their deck that are at stake.
// The battle card is the next one drawn after that.
// This event recursively continues until there is a winner.
// The winner at the end of the battle receives all cards at stake.
func war(d1 *Deck, d2 *Deck, d1Pile []Card, d2Pile []Card) {
	// 3 cards are at stake for each war.
	for i := 0; i < 3; i++ {
		if GameOver(d1, d2) {
			return
		}
		c1, _ := Draw(d1)
		c2, _ := Draw(d2)
		d1Pile = append(d1Pile, c1)
		d2Pile = append(d2Pile, c2)
	}
	// The battle determining card.
	readyToAdvance := false
	for !readyToAdvance {
		readyToAdvance = promptPlayer("A War is in progress. Are you view the results? (y/n) ")
	}
	print("War in progress")
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		print(".")
		if i == 2 {
			println()
		}
	}
	if GameOver(d1, d2) {
		return
	}
	c1, _ := Draw(d1)
	c2, _ := Draw(d2)
	result := c1.BattleAgainst(c2)
	if result == -1 { // If you lose, you lose all cards to the other player.
		d2.PlayingCards = append(d2.PlayingCards, c1)
		d2.PlayingCards = append(d2.PlayingCards, c2)
		d2.PlayingCards = append(d2.PlayingCards, d1Pile...)
		d2.PlayingCards = append(d2.PlayingCards, d2Pile...)
		fmt.Println("You lost this War...")
	} else if result == 0 { // Another War commences. Higher stakes...
		fmt.Println("Another War has commenced!")
		war(d1, d2, d1Pile, d2Pile)
	} else { // If you win, you receive all cards on the field!
		d1.PlayingCards = append(d1.PlayingCards, c1)
		d1.PlayingCards = append(d1.PlayingCards, c2)
		d1.PlayingCards = append(d1.PlayingCards, d1Pile...)
		d1.PlayingCards = append(d1.PlayingCards, d2Pile...)
		fmt.Println("You won this War!")
	}
}

func CommenceRound(d1 *Deck, d2 *Deck) {
	for i := 0; i < 10; i++ {
		println() // do this to reduce clutter on the screen
	}
	c1, _ := Draw(d1)
	c2, _ := Draw(d2)
	result := c1.BattleAgainst(c2)
	if result == -1 {
		d2.PlayingCards = append(d2.PlayingCards, c1)
		d2.PlayingCards = append(d2.PlayingCards, c2)
		fmt.Println("You lost this battle...")
	} else if result == 0 { // war commences
		d1Pile := make([]Card, 0)
		d2Pile := make([]Card, 0)
		d1Pile = append(d1Pile, c1)
		d2Pile = append(d2Pile, c2)
		war(d1, d2, d1Pile, d2Pile)
	} else {
		d1.PlayingCards = append(d1.PlayingCards, c2)
		d1.PlayingCards = append(d1.PlayingCards, c1)
		fmt.Println("You won this battle!")
	}
}

// Input: deck pointer, deck pointer
// Checks if either of the decks have 0 cards and returns true if either one does.
func GameOver(player *Deck, bot *Deck) bool {
	return len(player.PlayingCards) == 0 || len(bot.PlayingCards) == 0
}
