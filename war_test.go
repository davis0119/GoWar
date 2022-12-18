package main

import (
	"testing"

	"./deck"
)

func TestDeckInit(t *testing.T) {
	d := deck.DeckInit()
	if len(d.PlayingCards) != 52 {
		t.FailNow()
	}
}

func TestDeckDraw(t *testing.T) {
	d := deck.DeckInit()
	deck.Draw(d)
	if len(d.PlayingCards) != 51 {
		t.FailNow()
	}
	deck.Draw(d)
	deck.Draw(d)
	if len(d.PlayingCards) != 49 {
		t.FailNow()
	}
}

func TestDeckSplit(t *testing.T) {
	d := deck.DeckInit()
	d1, d2 := deck.SplitCards(d)
	if len(d1.PlayingCards) != 26 || len(d2.PlayingCards) != 26 {
		t.Log("Split Deck Failure")
		t.FailNow()
	}
	for _, c1 := range d1.PlayingCards {
		for _, c2 := range d2.PlayingCards {
			if c1.ToStr() == c2.ToStr() {
				t.Log("There should be no duplicate cards.")
				t.FailNow()
			}
		}
	}
}

func TestDeckSplitAndDraw(t *testing.T) {
	d := deck.DeckInit()
	d1, d2 := deck.SplitCards(d)
	deck.Draw(d1)
	if len(d1.PlayingCards) != 25 || len(d2.PlayingCards) != 26 {
		t.Log("Split Deck Failure")
		t.FailNow()
	}
}
