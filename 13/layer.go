package day13

// Position is function of time and range
func Position(t int, r int) int {
	if r == 1 {
		return 0
	}

	r = r - 1

	div := t / r
	mod := t % r

	// direction is down if division is even
	if div%2 == 0 {
		return mod
	} else {
		return r - mod
	}
}
