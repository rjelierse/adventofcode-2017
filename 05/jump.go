package trampoline

func CalculateSteps(input []int) (steps int) {
	var position int

	length := len(input)
	moves := make([]int, length)
	copy(moves, input)

	for position >= 0 && position < length {
		jump := moves[position]

		moves[position] = moves[position] + 1
		position = position + jump

		steps++
	}

	return steps
}

func CalculateSteps2(input []int) (steps int) {
	var position int

	length := len(input)
	moves := make([]int, length)
	copy(moves, input)

	for position >= 0 && position < length {
		jump := moves[position]

		if jump >= 3 {
			moves[position] = moves[position] - 1
		} else {
			moves[position] = moves[position] + 1
		}

		position = position + jump

		steps++
	}

	return steps
}
