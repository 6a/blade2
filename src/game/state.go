package game

import (
	"math/rand"
	"sort"
)

const maxDraws = 3

// State is a container for the current state of a game
type State struct {
	Turn  int
	Score [2]int
	Field [2][]CardID
	Hand  [2][]CardID
	Deck  [2][]CardID
}

// CreateGameState generates a new gamestate, containing all the objects and tracking values required for a single game
func CreateGameState() State {
	// Get all cards for this game
	stock := GenerateGameDeck()

	// filter out a subdeck for each player

	// Separate the deck into basics and specials
	basic := stock[0]
	specials := stock[1]

	// Determine how many specials each player will have (0-MaxSpecialCardsPerPlayer (3))
	p1NumSpecials := MinSpecialCardsPerPlayer + rand.Intn((MaxSpecialCardsPerPlayer-MinSpecialCardsPerPlayer)+1)
	p2NumSpecials := MinSpecialCardsPerPlayer + rand.Intn((MaxSpecialCardsPerPlayer-MinSpecialCardsPerPlayer)+1)

	// Must perform copy operations and work on the new slice otherwise the underlying slices get modified
	var p1Deck, p2Deck []CardID
	p1Deck = append(p1Deck, basic[0:MaxCardsPerPlayer-p1NumSpecials]...)
	p1Deck = append(p1Deck, specials[0:p1NumSpecials]...)

	p2Deck = append(p2Deck, basic[MaxCardsPerPlayer-p1NumSpecials:(MaxCardsPerPlayer-p1NumSpecials)+(MaxCardsPerPlayer-p2NumSpecials)]...)
	p2Deck = append(p2Deck, specials[p1NumSpecials:p1NumSpecials+p2NumSpecials]...)

	// Shuffle the cards again as currently any specials are stuck on the end of the container
	ShuffleCards(p1Deck)
	ShuffleCards(p2Deck)

	// Generate hands for each player
	h1, h2 := make([]CardID, HandSize), make([]CardID, HandSize)
	copy(h1, p1Deck[0:HandSize])
	copy(h2, p2Deck[0:HandSize])
	sort.Slice(h1, func(i, j int) bool { return int(h1[i]) < int(h1[j]) })
	sort.Slice(h2, func(i, j int) bool { return int(h2[i]) < int(h2[j]) })
	hand := [2][]CardID{h1, h2}

	// Each players deck is the remaining 6 cards in their subdeck minus 1 that is added to the field
	deck := [2][]CardID{p1Deck[HandSize : MaxCardsPerPlayer-1], p2Deck[HandSize : MaxCardsPerPlayer-1]}

	// The field is initialized with the last card in the deck that was leftover
	field := [2][]CardID{p1Deck[MaxCardsPerPlayer-1 : MaxCardsPerPlayer], p2Deck[MaxCardsPerPlayer-1 : MaxCardsPerPlayer]}

	// Both players start with the score from their first draw
	score := [2]int{field[0][0].Score(), field[1][0].Score()}

	// If the score is a draw, the turn needs to be calculated by looking forwards. If the next 3 draws result
	// in a draw, the cards are redrawn
	// Otherwise, the lowest score player goes first
	draws := 1
	s1, s2 := field[0][0].Score(), field[1][0].Score()
	turn := -1
	for {
		if s1 != s2 {
			if s1 > s2 {
				turn = 1
			} else {
				turn = 0
			}
			break
		}

		if draws >= maxDraws {
			return CreateGameState()
		}

		s1 = deck[0][len(deck[0])-draws].Score()
		s2 = deck[1][len(deck[1])-draws].Score()
		draws++
	}

	// Add all the precomputed values into a new State struct
	return State{turn, score, field, hand, deck}
}
