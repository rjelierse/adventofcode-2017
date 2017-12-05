package checksum

import (
	"log"
	"strconv"
	"strings"
)

func convertLine(input string) []int {
	input = strings.TrimSpace(input)
	numbers := strings.Split(input, "\t")

	return convertToInt(numbers)
}

func convertToInt(input []string) (output []int) {
	output = make([]int, len(input))

	for i, value := range input {
		number, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			log.Fatal("Cannot parse number:", err)
		}

		output[i] = int(number)
	}

	return
}
