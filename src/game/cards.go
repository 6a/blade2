package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/0110101001110011/blade2/src/e"
)

// Card represents a single card, with its name, description etc.
type Card struct {
	ID          int
	Title       string
	SpecialName string
	Description string
	Value       int
	Effect      e.Effect
}

func createSpecificCard(card e.CardID) Card {
	switch card {
	case e.ElliotStaff:
		return Card{int(card), "Elliot's Orbal Staff", "Rod", "Enable a card that was disabled by Bolt", 1, e.Rod}
	case e.Gunswords:
		return Card{int(card), "Fie's Twin Gunswords", "", "", 2, e.None}
	case e.Bow:
		return Card{int(card), "Alisa's Orbal Bow", "", "", 3, e.None}
	case e.Sword:
		return Card{int(card), "Jusis' Sword", "", "", 4, e.None}
	case e.Shotgun:
		return Card{int(card), "Machias' Orbal Shotgun", "", "", 5, e.None}
	case e.Spear:
		return Card{int(card), "Gaius' Spear", "", "", 6, e.None}
	case e.Greatsword:
		return Card{int(card), "Laura's Greatsword", "", "", 7, e.None}
	case e.Tachi:
		return Card{int(card), "Reans' Tachi", "Bolt", "Disables the last card placed by your oponent, and is then discarded", 1, e.Bolt}
	case e.EmmaStaff:
		return Card{int(card), "Emma's Orbal Staff", "Mirror", "Reverses the playing field, and is then discarded", 1, e.Mirror}
	case e.Rapier:
		return Card{int(card), "Elise's Rapier", "Blast", "Allows you to remove and discard a card from your oponents hand, and is then discarded", 1, e.Blast}
	case e.SwordAndGun:
		return Card{int(card), "Sara's Sword and Gun", "Force", "Doubles your score, and is then placed on the field", 1, e.Force}
	default:
		panic(fmt.Sprintf("A card with the ID [%T] was requested, but this is not a valid card ID", card))
	}
}

func createRandomCard() Card {
	rand.Seed(time.Now().UnixNano())
	cardID := e.CardID(rand.Intn(11))
	return createSpecificCard(cardID)
}

// CreateDeck generates a random deck of 30 cards to be used for a single game
func CreateDeck() [30]Card {
	deck := [30]Card{}

	// TODO magic

	return deck
}
