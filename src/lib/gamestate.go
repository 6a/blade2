package lib

type GameState struct {
	Turn  int
	Score [2]int
	Hand  [2][]Card
	Field [2][]Card
	Deck  [2][]Card
}

func NewGameState() GameState {
	cardSets := [2][]Card{}
	return GameState{0, [2]int{0, 0}, cardSets, cardSets, cardSets}
}
