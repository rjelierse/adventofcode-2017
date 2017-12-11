package garbage

// ReadStream parses a stream and returns the total group score
func ReadStream(input string) (score int, garbage int) {
	l := len(input)
	isGarbage := false
	level := 0

	for i := 0; i < l; i++ {
		// Escape character, ignore next and proceed
		if isGarbage && input[i] == '!' {
			i++
			continue
		}

		// Garbage sequence
		if isGarbage && input[i] != '>' {
			garbage++
			continue
		}

		switch input[i] {
		case '<':
			isGarbage = true
		case '>':
			isGarbage = false
		case '{':
			level++
			score = score + level
		case '}':
			level--
		}
	}

	return
}
