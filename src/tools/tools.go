package tools

// MaxInt returns the larger of two ints
func MaxInt(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

// AddRotaryInt adds two integers together. If the total exceeds, or is equal to <max>, <floor> is returned instead.
// For example, AddRotaryInt(4, 5, 10, 0) returns 9, while AddRotaryInt(5, 5, 10, 0) returns 0
func AddRotaryInt(a int, b int, max int, floor int) int {
	sum := a + b
	if sum >= max {
		return floor
	}

	return sum
}
