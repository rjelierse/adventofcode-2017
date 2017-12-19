package day16

type Instruction interface {
	Apply(*Floor)
}
