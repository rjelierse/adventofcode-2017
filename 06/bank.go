package day06


type Bank struct {
	Configuration          Configuration
	NumSlots               int
	PreviousConfigurations []Configuration
}

func (b *Bank) Redistribute() {
	configuration := b.Configuration.Clone()

	slot, value := configuration.High()
	configuration[slot] = 0

	for value > 0 {
		slot = (slot + 1) % b.NumSlots
		configuration[slot]++
		value--
	}

	b.PreviousConfigurations = append(b.PreviousConfigurations, b.Configuration)
	b.Configuration = configuration
}

func (b *Bank) ConfigurationSeenBefore() (bool, int) {
	for i, c := range b.PreviousConfigurations {
		if c.Is(b.Configuration) {
			return true, i
		}
	}

	return false, 0
}

func NewBank(configuration Configuration) *Bank {
	bank := new(Bank)

	bank.NumSlots = len(configuration)
	bank.Configuration = configuration.Clone()

	return bank
}
