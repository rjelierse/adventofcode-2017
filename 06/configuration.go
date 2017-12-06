package banks

type Configuration []int

func (c Configuration) High() (index int, value int) {
	for i, v := range c {
		if v > value {
			index = i
			value = v
		}
	}

	return
}

func (c Configuration) Is(t Configuration) bool {
	if len(c) != len(t) {
		return false
	}

	for i := 0; i < len(c); i++ {
		if c[i] != t[i] {
			return false
		}
	}

	return true
}

func (c Configuration) Clone() Configuration {
	return append(Configuration{}, c...)
}
