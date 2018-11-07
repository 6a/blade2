package templates

import (
	"encoding/json"

	"github.com/0110101001110011/blade2/src/e"
)

// GenericError is a prepackaged binary JSON used for unhandled error output
var GenericError = MakeJSON(Information{Code: e.UnknownError, Message: "Unhandled exception occured"})

// Information represents the outcome of a particular action
type Information struct {
	Code    e.Status
	Message string
}

// Delta represents a single change that occurs in a game
type Delta struct {
	Entity e.Entity
	Change int
	Index  int
}

// StateUpdate represents all the changes (delta) made during a players turn
type StateUpdate struct {
	TurnNumber int
	NextTurn   int // 0 or 1, and -1 for game ending
	Deltas     []Delta
}

// MakeJSON Packages struct as as binary JSON
func MakeJSON(data interface{}) []byte {
	res, _ := json.Marshal(data)
	return res
}
