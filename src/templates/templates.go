package templates

import (
	"encoding/json"

	"github.com/0110101001110011/blade2/src/e"
)

var GenericError = MakeJSON(Information{Status: e.UnknownError, Message: "Unhandled exception occured"})

type Information struct {
	Status  e.Status
	Message string
}

type Instruction struct {
	Instruction    e.Instruction
	MetadataLength int
	Metadata       []int
}

type Heartbeat struct {
	Uptime int64
}

func MakeJSON(data interface{}) []byte {
	res, _ := json.Marshal(data)
	return res
}
