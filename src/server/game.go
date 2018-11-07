package server

import (
	"github.com/0110101001110011/blade2/src/templates"
)

var nextGameID uint64

// Game is a container for a pair of clients and the associated gamestate
type Game struct {
	ID         uint64
	Client     [2]*Client
	StartState Cards
	History    []templates.StateUpdate
}

// CreateGame creates a new Game object using two Client objects
func CreateGame(c1 *Client, c2 *Client) Game {
	state := GenerateCards()
	game := Game{nextGameID, [2]*Client{c1, c2}, state, nil}
	c1.sendMessage(templates.MakeJSON(state))
	c2.sendMessage(templates.MakeJSON(state))
	nextGameID++
	return game
}

// RelayUpdates relays updates between clients, and ends the game if an endgame condition is detected
func (g *Game) RelayUpdates() {
	for clientIndex := range g.Client {
		if len(g.Client[clientIndex].Updates) > 0 {
			update := g.Client[clientIndex].Updates[0]
			g.Client[clientIndex].Updates = g.Client[clientIndex].Updates[1:]
			g.Client[1-clientIndex].sendMessage(templates.MakeJSON(update))
			g.History = append(g.History, update)
			if update.NextTurn == -1 {
				g.finish()
			}
		}
	}
}

func (g *Game) finish() {
	// TODO implement logic to sort history and dump to DB
}
