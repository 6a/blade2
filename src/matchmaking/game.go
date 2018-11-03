package matchmaking

import (
	"github.com/0110101001110011/blade2/src/game"
	"github.com/0110101001110011/blade2/src/templates"
	"github.com/0110101001110011/blade2/src/tools"
)

var nextGameID uint64

// Game is a container for a pair of clients and the associated gamestate
type Game struct {
	ID     uint64
	Client [2]*Client
	Turn   int
}

// CreateGame creates a new Game object using two Client objects
func CreateGame(c1 *Client, c2 *Client) Game {
	state := game.CreateGameState()
	game := Game{nextGameID, [2]*Client{c1, c2}, state.Turn}
	c1.sendMessage(templates.MakeJSON(state))
	c2.sendMessage(templates.MakeJSON(state))
	nextGameID++
	return game
}

// Update updates the internal game state send by a player and then relays it to the other player
func (g *Game) Update() {
	if g.Client[g.Turn].Update != nil {
		g.Client[tools.AddRotaryInt(g.Turn, 1, 2, 0)].sendMessage(templates.MakeJSON(g.Client[g.Turn].Update))

		if g.Client[g.Turn].Update.TurnChanged {
			g.Turn = tools.AddRotaryInt(g.Turn, 1, 2, 0)
		}

		g.Client[g.Turn].Update = nil
	}
}
