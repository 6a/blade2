package e

type Effect int

const (
	None   Effect = 0
	Rod    Effect = 1
	Bolt   Effect = 2
	Mirror Effect = 3
	Blast  Effect = 4
	Force  Effect = 5
)

type CardID int

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

type Status int

const (
	OK                       Status = 0
	Connected                Status = 1
	MatchFound               Status = 2
	UnknownError             Status = 3
	OponentDroppedConnection Status = 4
)

type Instruction int

const (
	GameFound Instruction = 0
)
