package game

// State is a container for the current state of a game
type State struct {
	Turn  int
	Score [2]int
	Hand  [2][]Card
	Field [2][]Card
	Deck  [2][]Card
}

// CreateGameState generates a new gamestate, containing all the objects and tracking values required for a single game
func CreateGameState() State {
	// Make Deck

	// Draw for each player

	// Create empty Field
	cardSets := [2][]Card{}
	return State{0, [2]int{0, 0}, cardSets, cardSets, cardSets}
}
