package tools

import (
	"math/rand"

	"github.com/0110101001110011/blade2/src/e"
)

// MaxInt returns the larger of two ints
func MaxInt(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

// AddRotaryInt adds two integers together. If the total exceed, or is equal to <max>, <floor> is returned instead
// For example:
// AddRotaryInt(4, 5, 10, 0) returns 9
// AddRotaryInt(5, 5, 10, 0) returns 0
// AddRotaryInt(7, 5, 10, 0) returns 0
// AddRotaryInt(7, 5, 10, 4) returns 4
func AddRotaryInt(a int, b int, max int, floor int) int {
	sum := a + b
	if sum >= max {
		return floor
	}

	return sum
}

func ShuffleCards(slice []e.CardID) {
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}
