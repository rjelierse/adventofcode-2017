package day16

type Floor struct {
	positions []byte
}

func NewFloor(positions int) *Floor {
	floor := new(Floor)
	floor.positions = make([]byte, 0, positions)

	start := byte('a')
	last := byte(int(start) + positions)
	for i := start; i != last; i++ {
		floor.positions = append(floor.positions, i)
	}

	return floor
}

func (floor *Floor) Dance(instructions []Instruction) string {
	for _, instruction := range instructions {
		instruction.Apply(floor)
	}

	return string(floor.positions)
}

func (floor *Floor) Spin(count int) {
	offset := len(floor.positions) - count
	positions := make([]byte, 0, len(floor.positions))

	positions = append(positions, floor.positions[offset:]...)
	positions = append(positions, floor.positions[:offset]...)

	floor.positions = positions
}

func (floor *Floor) Exchange(a int, b int) {
	charA := floor.positions[a]
	charB := floor.positions[b]

	floor.positions[a] = charB
	floor.positions[b] = charA
}

func (floor *Floor) Partner(a byte, b byte) {
	floor.Exchange(floor.getPosition(a), floor.getPosition(b))
}

func (floor *Floor) getPosition(char byte) int {
	for i, b := range floor.positions {
		if b == char {
			return i
		}
	}

	return -1
}
