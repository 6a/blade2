package matchmaking

import (
	"time"

	"github.com/0110101001110011/blade2/src/game"
	"github.com/0110101001110011/blade2/src/templates"
)

var nextMatchID = 0

type Match struct {
	ID     int
	Client [2]Client
	State  game.GameState
}

func CreateMatch(c1 Client, c2 Client) Match {
	match := Match{nextMatchID, [2]Client{c1, c2}, game.CreateGameState()}
	nextMatchID++
	return match
}

func (m *Match) ExecuteNext() {
	time.Sleep(1 * time.Second)
	m.Client[m.State.Turn].sendMessage(templates.Make(templates.StandardJSON{Status: game.OK, Message: "Did something"}))
	m.ChangeTurn()
}

func (m *Match) ChangeTurn() {
	currentTurn := m.State.Turn
	if currentTurn == 0 {
		m.State.Turn = 1
	} else {
		m.State.Turn = 0
	}
}
