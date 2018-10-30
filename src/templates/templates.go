package templates

import (
	"encoding/json"

	"github.com/0110101001110011/blade2/src/game"
)

var GenericError, err = json.Marshal(StandardJSON{game.UnknownError, "Unhandled exception occured"})

type StandardJSON struct {
	Status  game.EStatus
	Message string
}

type Instruction struct {
	Instruction game.EInstruction
}

type Heartbeat struct {
	Uptime int64
}
