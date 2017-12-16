package dance

import (
	"fmt"
)

type ExchangeInstruction struct {
	PositionA int
	PositionB int
}

func (i *ExchangeInstruction) Apply(d *Floor) {
	d.Exchange(i.PositionA, i.PositionB)
}

func ParseExchangeInstruction(instruction string) (*ExchangeInstruction, error) {
	i := new(ExchangeInstruction)
	if _, err := fmt.Sscanf(instruction, "x%d/%d", &i.PositionA, &i.PositionB); err != nil {
		return nil, err
	}
	return i, nil
}
