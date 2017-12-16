package dance

import (
	"fmt"
)

type SpinInstruction struct {
	Count int
}

func (i *SpinInstruction) Apply(d *Dancefloor) {
	d.Spin(i.Count)
}

func ParseSpinInstruction(instruction string) (*SpinInstruction, error) {
	i := new(SpinInstruction)
	if _, err := fmt.Sscanf(instruction, "s%d", &i.Count); err != nil {
		return nil, err
	}
	return i, nil
}
