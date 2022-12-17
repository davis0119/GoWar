package deck

import (
	"math/rand"
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

func (c *Card) String() string {
	return ""
}

func (c1 *Card) BattleAgainst(c2 Card) int {
	if c1.Value.Val < c2.Value.Val {
		return -1 // loses
	} else if c1.Value.Val == c2.Value.Val {
		return 0 // tie (WAR)
	} else {
		return 1 // win
	}
}

func Shuffle(d Deck) Deck {
	for i := 1; i < len(d.PlayingCards); i++ {
		randomNum := rand.Intn(i + 1)
		if i != randomNum { // swap the cards
			d.PlayingCards[randomNum], d.PlayingCards[i] = d.PlayingCards[i], d.PlayingCards[randomNum]
		}
	}
	return d
}

func Draw(d *Deck) (Card, *Deck) {
	card := d.PlayingCards[0]
	d.PlayingCards = d.PlayingCards[1:]
	return card, d
}

func CommenceRound(d1 *Deck, d2 *Deck) {
	c1, d1 := Draw(d1)
	c2, d2 := Draw(d2)
	result := c1.BattleAgainst(c2)
	if result == -1 {
		d2.PlayingCards = append(d2.PlayingCards, c1)
		d2.PlayingCards = append(d2.PlayingCards, c2)
	} else if result == 0 { // war commencess

	} else {
		d1.PlayingCards = append(d1.PlayingCards, c2)
		d1.PlayingCards = append(d1.PlayingCards, c1)
	}
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
		{"Two", 2},
		{"Three", 3},
		{"Four", 4},
		{"Five", 5},
		{"Six", 6},
		{"Seven", 7},
		{"Eight", 8},
		{"Nine", 9},
		{"Ten", 10},
		{"Jack", 11},
		{"Queen", 12},
		{"King", 13},
		{"Ace", 14},
	}
	d.PlayingCards = nil
	for _, v := range values {
		for _, s := range suits {
			d.PlayingCards = append(d.PlayingCards, Card{Suit: s, Value: v})
		}
	}
	return d
}
