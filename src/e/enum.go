package e

// Effect represents different effects that a card can have
type Effect int

// Effect consts (enum)
const (
	None   Effect = 0
	Rod    Effect = 1
	Bolt   Effect = 2
	Mirror Effect = 3
	Blast  Effect = 4
	Force  Effect = 5
)

// GenerateDeck represents the ID number for each card
type CardID int

// CardID consts (enum)
const (
	ElliotStaff CardID = 0
	Gunswords   CardID = 1
	Bow         CardID = 2
	Sword       CardID = 3
	Shotgun     CardID = 4
	Spear       CardID = 5
	Greatsword  CardID = 6
	Tachi       CardID = 7
	EmmaStaff   CardID = 8
	Rapier      CardID = 9
	SwordAndGun CardID = 10
)

// PayloadType represents various status codes for communication between the client and servers
type PayloadType int

// Status consts (enum) (OK or neutral)
const (
	OK         PayloadType = 100
	Connected  PayloadType = 101
	MatchFound PayloadType = 102
)

// Status consts (enum) (errors)
const (
	UnknownError             PayloadType = 200
	OponentDroppedConnection PayloadType = 201
)

// Instruction represents an instruction code, used by a client to determine what to do with an instruction payload
type Instruction int

// Instruction consts (enum)
const (
	GameFound Instruction = 0
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
