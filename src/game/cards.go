package game

import (
	"math/rand"
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

// GenerateGameDeck returns an random array of cards to be used for a single game
func GenerateGameDeck() [2][]CardID {
	// Make a copy of the base deck and shuffle them
	basic := make([]CardID, baseDeckBasicSize)
	copy(basic, baseDeckBasic)
	ShuffleCards(basic)

	// Make a copy of the specials deck and shuffle them
	special := make([]CardID, baseDeckSpecialSize)
	copy(special, baseDeckSpecial)
	ShuffleCards(special)

	outdecks := [2][]CardID{basic[0:baseDeckBasicSize], special[0:maxSpecials]}

	return outdecks
}

// Score returns the basic score value for a card as if it was played without triggering any effect
func (c CardID) Score() int {
	raw := int(c)
	if raw <= 6 {
		return raw + 1
	}

	return 1
}

func ShuffleCards(slice []CardID) {
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}
