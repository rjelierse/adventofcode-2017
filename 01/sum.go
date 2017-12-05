package captcha

const AsciiOffset = 48

func calculateSum(input []byte, offset int) (sum int) {
	length := len(input)

	for i := 0; i < length; i++ {
		j := (i + offset) % length

		a := int(input[i]) - AsciiOffset
		b := int(input[j]) - AsciiOffset

		if a == b {
			sum = sum + a
		}
	}

	return sum
}
