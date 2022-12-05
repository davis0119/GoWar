package deck

type Face String
type Suit String
type Value int

type Card struct {
	Face  Face
	Suit  Suit
	Value Value
}

type Deck struct {
	PlayingCards []Card
}
