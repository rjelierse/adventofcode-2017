package day19

type Node byte

func (n Node) IsVoid() bool {
	return n == ' '
}

func (n Node) IsNotVoid() bool {
	return n != ' '
}

func (n Node) IsTurn() bool {
	return n == '+'
}

func (n Node) IsAlpha() bool {
	return (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'Z')
}
