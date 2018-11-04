package matchmaking

import (
	"github.com/0110101001110011/blade2/src/game"
	"github.com/0110101001110011/blade2/src/templates"
)

var nextGameID uint64

// Game is a container for a pair of clients and the associated gamestate
type Game struct {
	ID     uint64
	Client [2]*Client
}

// CreateGame creates a new Game object using two Client objects
func CreateGame(c1 *Client, c2 *Client) Game {
	state := game.GenerateCards()
	game := Game{nextGameID, [2]*Client{c1, c2}}
	c1.sendMessage(templates.MakeJSON(state))
	c2.sendMessage(templates.MakeJSON(state))
	nextGameID++
	return game
}

// RelayUpdates relays any updates from a client to the other client
func (g *Game) RelayUpdates() {
	if g.Client[0].Update != nil {
		g.Client[1].sendMessage(templates.MakeJSON(g.Client[0].Update))
		g.Client[0].Update = nil
	}

	if g.Client[1].Update != nil {
		g.Client[0].sendMessage(templates.MakeJSON(g.Client[1].Update))
		g.Client[1].Update = nil
	}
}
