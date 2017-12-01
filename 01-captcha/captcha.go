package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const AsciiOffset = 48

func main() {
	input := readInput(os.Args[1])
	// Part 1:
	fmt.Println("[1] Captcha result:", calculateSum(input, 1))

	// Part 2:
	fmt.Println("[2] Captcha result:", calculateSum(input, len(input)/2))
}

func readInput(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Cannot read file:", err)
	}

	return bytes
}

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
