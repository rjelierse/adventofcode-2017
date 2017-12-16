package dance

type PartnerInstruction struct {
	PartnerA byte
	PartnerB byte
}

func (i *PartnerInstruction) Apply(d *Dancefloor) {
	d.Partner(i.PartnerA, i.PartnerB)
}

func ParsePartnerInstruction(instruction string) *PartnerInstruction {
	i := new(PartnerInstruction)
	i.PartnerA = instruction[1]
	i.PartnerB = instruction[3]
	return i
}
