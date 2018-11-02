package game

import (
	"math/rand"

	"github.com/0110101001110011/blade2/src/e"
)

// State is a container for the current state of a game
type State struct {
	Turn  int
	Score [2]int
	Hand  [2][]e.CardID
	Field [2][]e.CardID
	Deck  [2][]e.CardID
}

// CreateGameState generates a new gamestate, containing all the objects and tracking values required for a single game
func CreateGameState() State {
	// Get all cards for this game
	stock := GenerateGameDeck()

	// Draw each players hand. As the cards are preshuffled, they are simply taken in sequence
	hand := [2][]e.CardID{stock[0:10], stock[10:20]}

	// Each players deck is half of the 10 remaining cards
	deck := [2][]e.CardID{stock[20:25], stock[25:30]}

	// The field starts as an empty array of cards
	field := [2][]e.CardID{}

	// Both players start at 0 score
	score := [2]int{0, 0}

	// Randomly select who goes first
	turn := rand.Intn(2)

	// Add all the precomputed values into a new State struct
	return State{turn, score, hand, field, deck}
}
