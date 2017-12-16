package dance

type Instruction interface {
	Apply(*Dancefloor)
}
