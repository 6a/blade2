package e

// Status represents various status codes for communication between the client and servers
type Status int

// Status consts (enum) (OK or neutral)
const (
	OK         Status = 100
	Connected  Status = 101
	MatchFound Status = 102
)

// Status consts (enum) (errors)
const (
	UnknownError             Status = 200
	OponentDroppedConnection Status = 201
	Drop                     Status = 202
)

// Entity describes a changeable part of the game (hand, deck etc)
type Entity int

// Entity consts (enum)
const (
	Hand         Entity = 0
	Deck         Entity = 1
	PlayerField  Entity = 2
	OponentField Entity = 3
	Score        Entity = 4
	FieldSwap    Entity = 5
	ClearField   Entity = 6
)
