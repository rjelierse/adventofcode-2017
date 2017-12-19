package day19

func Travel(grid Grid) (sequence string, steps int) {
	position := grid.Start()
	direction := DirectionDown
	steps = 1

	for {
		position = position.Move(direction)
		if !grid.InGrid(position) {
			return
		}

		node := grid.Get(position)
		switch {
		case node.IsVoid():
			return
		case node.IsTurn():
			direction = direction.Change(grid, position)
		case node.IsAlpha():
			sequence += string(node)
		}

		steps++

		if direction == DirectionNone {
			return
		}
	}
}
