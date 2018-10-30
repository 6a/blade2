package game

type EEffect int

const (
	Unbolt EEffect = 0
	Bolt   EEffect = 1
	Mirror EEffect = 2
	Blast  EEffect = 3
	Force  EEffect = 4
)

type EStatus int

const (
	OK           EStatus = 0
	UnknownError EStatus = 1
)

type EInstruction int

const (
	GameFound EInstruction = 0
)
