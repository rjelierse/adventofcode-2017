package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	input := strings.Split(readInput(os.Args[1]), "")
	// Part 1:
	fmt.Println("[1] Captcha result:", calculateSum(input, 1))

	// Part 2:
	fmt.Println("[2] Captcha result:", calculateSum(input, len(input) / 2))
}

func readInput(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Cannot read file:", err)
	}

	return string(bytes[:])
}

func calculateSum(input []string, offset int) (sum int64) {
	length := len(input)

	for i := 0; i < length; i++ {
		j := (i + offset) % length
		current, err := strconv.ParseInt(input[i], 10, 0)
		if err != nil {
			log.Fatal("Cannot parse number:", err)
		}

		next, err := strconv.ParseInt(input[j], 10, 0)
		if err != nil {
			log.Fatal("Cannot parse number:", err)
		}

		if current == next {
			sum = sum + current
		}
	}

	return sum
}
