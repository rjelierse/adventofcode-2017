package day02

import (
	"sort"
)

// Difference calculates the distance between the highest and lowest value in a slice
func Difference(numbers []int) int {
	sort.Ints(numbers)

	return numbers[len(numbers)-1] - numbers[0]
}

// Division calculates the division between the two values that can divide in a slice
func Division(numbers []int) int {
	sort.Ints(numbers)

	length := len(numbers)

	for i, a := range numbers {
		for rev := length - 1; rev > i; rev-- {
			b := numbers[rev]

			if b%a == 0 {
				return b / a
			}
		}
	}

	return 0
}
