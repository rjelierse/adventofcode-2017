package dance

type Dancefloor struct {
	Positions []byte
	history   []string
}

func NewDancefloor(positions int) *Dancefloor {
	d := new(Dancefloor)
	d.Positions = make([]byte, 0, positions)
	d.history = make([]string, 0)

	start := byte('a')
	last := byte(int(start) + positions)
	for i := start; i != last; i++ {
		d.Positions = append(d.Positions, i)
	}

	return d
}

func (d *Dancefloor) Dance(instructions []string) {
	for _, instruction := range instructions {
		switch GetInstruction(instruction) {
		case InstructionSpin:
			count := ParseSpinInstruction(instruction)
			d.Spin(count)
		case InstructionExchange:
			a, b := ParseExchangeInstruction(instruction)
			d.Exchange(a, b)
		case InstructionPartner:
			a, b := ParsePartnerInstruction(instruction)
			d.Partner(a, b)
		}
	}

	d.history = append(d.history, string(d.Positions))
}

func (d *Dancefloor) IsRepeat() int {
	position := string(d.Positions)

	for i, p := range d.history {
		if p == position {
			return i
		}
	}

	return -1
}

func (d *Dancefloor) Spin(count int) {
	offset := len(d.Positions) - count
	positions := make([]byte, 0, len(d.Positions))

	positions = append(positions, d.Positions[offset:]...)
	positions = append(positions, d.Positions[:offset]...)

	d.Positions = positions
}

func (d *Dancefloor) Exchange(a int, b int) {
	charA := d.Positions[a]
	charB := d.Positions[b]

	d.Positions[a] = charB
	d.Positions[b] = charA
}

func (d *Dancefloor) Partner(a byte, b byte) {
	d.Exchange(d.getPosition(a), d.getPosition(b))
}

func (d *Dancefloor) getPosition(char byte) int {
	for i, b := range d.Positions {
		if b == char {
			return i
		}
	}

	return -1
}
