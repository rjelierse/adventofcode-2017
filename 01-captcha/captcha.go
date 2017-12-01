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
	sum := int64(0)
	input = append(input, input[0])

	for i := 0; i < (len(input) - 1); i++ {
		current, err := strconv.ParseInt(input[i], 10, 0)
		if err != nil {
			log.Fatal("Cannot parse number:", err)
		}

		next, err := strconv.ParseInt(input[i+1], 10, 0)
		if err != nil {
			log.Fatal("Cannot parse number:", err)
		}

		if current == next {
			sum = sum + current
		}
	}

	fmt.Println("Sum of adjacent numbers:", sum)
}

func readInput(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Cannot read file:", err)
	}

	return string(bytes[:])
}
