package templates

import (
	"encoding/json"

	"github.com/0110101001110011/blade2/src/e"
)

var GenericError = MakeJSON(Information{Code: e.UnknownError, Message: "Unhandled exception occured"})

type Information struct {
	Code    e.PayloadType
	Message string
}

type Instruction struct {
	Code     e.PayloadType
	Metadata []int
}

type Delta struct {
	Entity e.Entity
	Change int
	Index  int
}

// StateUpdate represents all the changes (delta) made during a players turn
type StateUpdate struct {
	Deltas []Delta
}

func MakeJSON(data interface{}) []byte {
	res, _ := json.Marshal(data)
	return res
}
