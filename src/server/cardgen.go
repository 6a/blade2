package server

import (
	"math/rand"
	"sort"
)

// MinSpecialCardsPerPlayer is the minimum number of special cards that each player can have in a single game
const MinSpecialCardsPerPlayer = 3

// MaxSpecialCardsPerPlayer is the maximum number of special cards that each player can have in a single game
const MaxSpecialCardsPerPlayer = 4

// MaxCardsPerPlayer is the number of cards that a single player has in a single game (this is the hand + deck count at the start of a game)
const MaxCardsPerPlayer = 16

// HandSize is the number of cards that a player starts with in their hand
const HandSize = 10

// PlayerDeckSize is the number of cards that a player starts with in their personal deck
const PlayerDeckSize = 6

const deckSize = MaxCardsPerPlayer * 2
const maxSpecials = MaxSpecialCardsPerPlayer * 3
const maxDraws = 3

var baseDeckSpecial = makeBaseDeckSpecials()
var baseDeckBasic = makeBaseDeckBasic()
var baseDeckBasicSize = len(baseDeckBasic)
var baseDeckSpecialSize = len(baseDeckSpecial)

// CardID represents the ID number for each card
type CardID int

// CardID consts (enum)
const (
	ElliotStaff CardID = 0
	Gunswords   CardID = 1
	Bow         CardID = 2
	Sword       CardID = 3
	Shotgun     CardID = 4
	Spear       CardID = 5
	Greatsword  CardID = 6
	Tachi       CardID = 7
	EmmaStaff   CardID = 8
	Rapier      CardID = 9
	SwordAndGun CardID = 10
)

func makeBaseDeckBasic() []CardID {
	// This implementation is a bit primitive but at least its easy to see whats going into the stock
	var deck []CardID
	deck = append(deck, ElliotStaff, ElliotStaff, ElliotStaff, ElliotStaff)
	deck = append(deck, Gunswords, Gunswords, Gunswords, Gunswords)
	deck = append(deck, Bow, Bow, Bow, Bow)
	deck = append(deck, Sword, Sword, Sword, Sword)
	deck = append(deck, Shotgun, Shotgun, Shotgun, Shotgun)
	deck = append(deck, Spear, Spear, Spear, Spear)
	deck = append(deck, Greatsword, Greatsword)

	return deck
}

func makeBaseDeckSpecials() []CardID {
	// This implementation is a bit primitive but at least its easy to see whats going into the stock
	var deck []CardID
	deck = append(deck, Tachi, Tachi, Tachi, Tachi)
	deck = append(deck, EmmaStaff, EmmaStaff, EmmaStaff, EmmaStaff)
	deck = append(deck, Rapier, Rapier, Rapier, Rapier)
	deck = append(deck, SwordAndGun, SwordAndGun)

	return deck
}

// Cards is a container for all the cards in a single game
type Cards struct {
	Field [2][]CardID
	Hand  [2][]CardID
	Deck  [2][]CardID
}

// GenerateCards generates a new Cards struct, containing all the objects and tracking values required for a single game
func GenerateCards() Cards {
	// Get all cards for this game
	basic, specials := GenerateGameDeck()

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

	// If the score is a draw, search forwards to ensure that the game will progress within the next 3 draws.
	// If the turn is undecided after 3 draws from each players deck, the cards are redrawn
	draws := 1
	s1, s2 := field[0][0].Score(), field[1][0].Score()
	for {
		if s1 != s2 {
			break
		}

		if draws >= maxDraws {
			return GenerateCards()
		}

		s1 = deck[0][len(deck[0])-draws].Score()
		s2 = deck[1][len(deck[1])-draws].Score()
		draws++
	}

	// Add all the precomputed values into a new Cards struct
	return Cards{field, hand, deck}
}

// GenerateGameDeck returns an random array of cards to be used for a single game
func GenerateGameDeck() (basic, special []CardID) {
	// Make a copy of the base deck and shuffle them
	b := make([]CardID, baseDeckBasicSize)
	copy(b, baseDeckBasic)
	ShuffleCards(b)

	// Make a copy of the specials deck and shuffle them
	s := make([]CardID, baseDeckSpecialSize)
	copy(s, baseDeckSpecial)
	ShuffleCards(s)

	return b[0:baseDeckBasicSize], s[0:maxSpecials]
}

// Score returns the basic score value for a card as if it was played without triggering any effect
func (c CardID) Score() int {
	raw := int(c)
	if raw <= 6 {
		return raw + 1
	}

	return 1
}

// ShuffleCards takes a slice of CardID's and randomly sorts them in-place (slices are passed by reference)
func ShuffleCards(slice []CardID) {
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}
