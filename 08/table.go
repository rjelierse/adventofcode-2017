package registers

var table = make(map[string]int)

func GetValue(register string) int {
	value, exists := table[register]
	if !exists {
		value = 0
	}

	return value
}

func SetValue(register string, value int) {
	table[register] = value
}

func LargestValue() int {
	value := 0

	for _, v := range table {
		if v > value {
			value = v
		}
	}

	return value
}
