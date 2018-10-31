package matchmaking

import (
	"time"

	"github.com/0110101001110011/blade2/src/e"
	"github.com/0110101001110011/blade2/src/gameobjects"
	"github.com/0110101001110011/blade2/src/templates"
)

var nextGameID uint64

// Game is a container for a pair of clients and the associated gamestate
type Game struct {
	ID     uint64
	Client [2]*Client
	State  gameobjects.GameState
}

// CreateGame creates a new Game object using two Client objects
func CreateGame(c1 *Client, c2 *Client) Game {
	game := Game{nextGameID, [2]*Client{c1, c2}, gameobjects.CreateGameState()}
	nextGameID++
	return game
}

// ExecuteNext prompts a Game object to advance the game by 1 step
func (g *Game) ExecuteNext() {
	time.Sleep(1 * time.Second)
	g.Client[g.State.Turn].sendMessage(templates.MakeJSON(templates.Information{Status: e.OK, Message: "Did something"}))
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
