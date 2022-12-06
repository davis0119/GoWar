package deck

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

func (d *Deck) DeckInit() {
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
}
