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

// Stringify the card.
func (c *Card) ToStr() string {
	return c.Value.Name + " of " + c.Suit.Name
}

// Shuffles the Deck of cards randomly.
func Shuffle(d *Deck) *Deck {
	rand.Seed(time.Now().UnixNano()) // not doing this doesn't really "shuffle" the deck every time
	for i := len(d.PlayingCards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.PlayingCards[i], d.PlayingCards[j] = d.PlayingCards[j], d.PlayingCards[i] // swap the cards
	}
	return d
}

// Draw a card from a deck. The card is removed from the deck.
func Draw(d *Deck) (Card, *Deck) {
	card := d.PlayingCards[0]
	d.PlayingCards = d.PlayingCards[1:]
	return card, d
}

// Split the deck in half. Used to give cards to both players.
func SplitCards(d *Deck) (player *Deck, bot *Deck) {
	player = new(Deck)
	bot = new(Deck)
	for len(d.PlayingCards) > 0 {
		c1, _ := Draw(d)
		c2, _ := Draw(d)
		player.PlayingCards = append(player.PlayingCards, c1)
		bot.PlayingCards = append(bot.PlayingCards, c2)
	}
	return player, bot
}

// Strigify the deck.
func (d *Deck) ToStr() string {
	s := ""
	for _, c := range d.PlayingCards {
		println(c.ToStr())
	}
	return s
}

// Used as a utility function to reduce clutter in the terminal.
func FocusTerminal() {
	for i := 0; i < 10; i++ {
		println()
	}
}

// Initialize the deck with a bunch of cards.
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

// Similar to a compareTo method. Sees which card would be victorious in battle.
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

// Awaits the player prompt to advance in the war.
func promptPlayer(prompt string) int {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scanln(&input)
		if len(input) > 0 {
			if input[0] == 'y' {
				return 1
			}
			if input[0] == 'n' {
				return 0
			}
			if input[0] == 'f' {
				return -1
			}
			if input == "help" {
				FocusTerminal()
				fmt.Println("War Event: Occurs due to both players drawing the same rank card.")
				fmt.Println("In this event, both players place down 3 cards and draw one additional card to battle.")
				fmt.Println("The winner takes all cards on the field.")
				fmt.Println("If there is another tie, this War Event recursively continues. If a player runs out of cards, they lose the game.")
			}
		}
		if input != "help" {
			fmt.Println("Answer via 'y' or 'n'. If you would like more info, type 'help'")
		}
	}
}

// The event when 2 cards are played with equal rank.
// Each player draws 3 cards from their deck that are at stake.
// The battle card is the next one drawn after that.
// This event recursively continues until there is a winner.
// The winner at the end of the battle receives all cards at stake.
func war(player *Deck, bot *Deck, playerPile []Card, botPile []Card) {
	// 3 cards are at stake for each war.
	for i := 0; i < 3; i++ {
		if GameOver(player, bot) {
			return
		}
		c1, _ := Draw(player)
		c2, _ := Draw(bot)
		playerPile = append(playerPile, c1)
		botPile = append(botPile, c2)
	}
	// The battle determining card.
	readyToAdvance := 0
	for readyToAdvance == 0 {
		readyToAdvance = promptPlayer("A War Event is in progress. Are you ready to view the results? (y/n | f to forfeit this War Event) ")
	}
	if GameOver(player, bot) {
		return
	}
	if readyToAdvance == -1 {
		warWaiting("You have forfeited this War")
		bot.PlayingCards = append(bot.PlayingCards, playerPile...)
		bot.PlayingCards = append(bot.PlayingCards, botPile...)
		fmt.Println("Your troops have now joined your opponent's side.")
		return
	}
	warWaiting("War in progress")
	c1, _ := Draw(player)
	c2, _ := Draw(bot)
	result := c1.BattleAgainst(c2)
	if result == -1 { // If you lose, you lose all cards to the other player.
		bot.PlayingCards = append(bot.PlayingCards, c1)
		bot.PlayingCards = append(bot.PlayingCards, c2)
		bot.PlayingCards = append(bot.PlayingCards, playerPile...)
		bot.PlayingCards = append(bot.PlayingCards, botPile...)
		fmt.Println("You lost this War...")
	} else if result == 0 { // Another War commences. Higher stakes...
		fmt.Println("Another War has commenced!")
		war(player, bot, playerPile, botPile)
	} else { // If you win, you receive all cards on the field!
		player.PlayingCards = append(player.PlayingCards, c1)
		player.PlayingCards = append(player.PlayingCards, c2)
		player.PlayingCards = append(player.PlayingCards, playerPile...)
		player.PlayingCards = append(player.PlayingCards, botPile...)
		fmt.Println("You won this War!")
	}
}

// Represents a battle / round of War. Players both draw a card and see which player wins the battle.
// The winner receives the cards on the field. The player with no cards remaining loses.
func CommenceRound(player *Deck, bot *Deck) {
	FocusTerminal()
	c1, _ := Draw(player)
	c2, _ := Draw(bot)
	result := c1.BattleAgainst(c2)
	if result == -1 {
		bot.PlayingCards = append(bot.PlayingCards, c1)
		bot.PlayingCards = append(bot.PlayingCards, c2)
		fmt.Println("You lost this battle...")
	} else if result == 0 { // war commences
		playerPile := make([]Card, 0)
		botPile := make([]Card, 0)
		playerPile = append(playerPile, c1)
		botPile = append(botPile, c2)
		war(player, bot, playerPile, botPile)
	} else {
		player.PlayingCards = append(player.PlayingCards, c2)
		player.PlayingCards = append(player.PlayingCards, c1)
		fmt.Println("You won this battle!")
	}
}

// Input: deck pointer, deck pointer
// Checks if either of the decks have 0 cards and returns true if either one does.
func GameOver(player *Deck, bot *Deck) bool {
	return len(player.PlayingCards) == 0 || len(bot.PlayingCards) == 0
}

// Used as a utility function to add effect to the game. As troops are battling, the players await in suspense.
func warWaiting(message string) {
	print(message)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		print(".")
		if i == 2 {
			println()
		}
	}
}
