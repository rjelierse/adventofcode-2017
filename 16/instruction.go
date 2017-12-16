package dance

type Instruction interface {
	Apply(*Floor)
}
