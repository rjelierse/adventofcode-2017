package day08

const (
	OperatorIncrease Operator = "inc"
	OperatorDecrease          = "dec"
)

type Operator string

type Instruction struct {
	Register string
	Operator Operator
	Value    int
}


func (i *Instruction) Apply() {
	value := GetValue(i.Register)

	switch i.Operator {
	case OperatorIncrease:
		value = value + i.Value
	case OperatorDecrease:
		value = value - i.Value
	}

	SetValue(i.Register, value)
}
