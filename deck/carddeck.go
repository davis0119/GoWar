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

func CommenceRound(d1 *Deck, d2 *Deck) {
	for i := 0; i < 10; i++ {
		println() // do this to reduce clutter on the screen
	}
	c1, d1 := Draw(d1)
	c2, d2 := Draw(d2)
	result := c1.BattleAgainst(c2)
	if result == -1 {
		d2.PlayingCards = append(d2.PlayingCards, c1)
		d2.PlayingCards = append(d2.PlayingCards, c2)
		fmt.Println("You lost this battle...")
	} else if result == 0 { // war commences
		fmt.Println("War has commenced!")
	} else {
		d1.PlayingCards = append(d1.PlayingCards, c2)
		d1.PlayingCards = append(d1.PlayingCards, c1)
		fmt.Println("You won this battle!")
	}
}
