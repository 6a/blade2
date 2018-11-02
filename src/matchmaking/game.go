package matchmaking

import (
	"time"

	"github.com/0110101001110011/blade2/src/e"
	"github.com/0110101001110011/blade2/src/game"
	"github.com/0110101001110011/blade2/src/templates"
)

var nextGameID uint64

// Game is a container for a pair of clients and the associated gamestate
type Game struct {
	ID      uint64
	Client  [2]*Client
	State   game.State
	Started bool
}

// CreateGame creates a new Game object using two Client objects
func CreateGame(c1 *Client, c2 *Client) Game {
	game := Game{nextGameID, [2]*Client{c1, c2}, game.CreateGameState(), false}
	nextGameID++
	return game
}

// Update updates the internal game state send by a player and then relays it to the other player
func (g *Game) Update() {
	if !g.Started {
		// TODO send gamestate to both players
		g.Started = true
		return
	}
	time.Sleep(1 * time.Second)
	g.Client[g.State.Turn].sendMessage(templates.MakeJSON(templates.Information{Code: e.OK, Message: "Did something"}))
	g.changeTurn()
}

func (g *Game) changeTurn() {
	currentTurn := g.State.Turn
	if currentTurn == 0 {
		g.State.Turn = 1
	} else {
		g.State.Turn = 0
	}
}
