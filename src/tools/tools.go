package tools

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func AddRotary(a int, b int, max int, floor int) int {
	sum := a + b
	if sum >= max {
		return floor
	} else {
		return sum
	}
}
