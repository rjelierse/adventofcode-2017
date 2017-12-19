package day02

import (
	"io/ioutil"
	"log"
)

func readInput(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Cannot read file:", err)
	}

	return string(bytes)
}
