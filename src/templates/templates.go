package templates

import (
	"encoding/json"

	"github.com/0110101001110011/blade2/src/game"
)

var GenericError = Make(StandardJSON{Status: game.UnknownError, Message: "Unhandled exception occured"})

type StandardJSON struct {
	Status  game.EStatus
	Message string
}

type Instruction struct {
	Instruction    game.EInstruction
	MetadataLength int
	Metadata       []int
}

type Heartbeat struct {
	Uptime int64
}

func Make(data interface{}) []byte {
	res, _ := json.Marshal(data)
	return res
}
