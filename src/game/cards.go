package game

import (
	"github.com/0110101001110011/blade2/src/e"
	"github.com/0110101001110011/blade2/src/tools"
)

const DeckSize = 30

// Returns an random array of cards to be used for a single game
func GenerateGameDeck() []e.CardID {
	// This implementation is a bit primitive but at least its easy to see whats going into the stock
	var stock []e.CardID
	stock = append(stock, e.ElliotStaff, e.ElliotStaff, e.ElliotStaff, e.ElliotStaff)
	stock = append(stock, e.Gunswords, e.Gunswords, e.Gunswords, e.Gunswords)
	stock = append(stock, e.Bow, e.Bow, e.Bow, e.Bow)
	stock = append(stock, e.Sword, e.Sword, e.Sword, e.Sword)
	stock = append(stock, e.Shotgun, e.Shotgun, e.Shotgun, e.Shotgun)
	stock = append(stock, e.Spear, e.Spear, e.Spear, e.Spear)
	stock = append(stock, e.Greatsword, e.Greatsword) // Note - Only 2 of these?
	stock = append(stock, e.Tachi, e.Tachi, e.Tachi, e.Tachi)
	stock = append(stock, e.EmmaStaff, e.EmmaStaff, e.EmmaStaff, e.EmmaStaff)
	stock = append(stock, e.Rapier, e.Rapier, e.Rapier, e.Rapier)
	stock = append(stock, e.SwordAndGun, e.SwordAndGun, e.SwordAndGun, e.SwordAndGun) // Note - Only 2 of these?

	// Shuffle all the cards
	tools.ShuffleCards(stock)

	// Take the first 30 cards and return them
	return stock[0:30]
}
