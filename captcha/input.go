package captcha

import (
	"io/ioutil"
	"log"
)

func readInput(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Cannot read file:", err)
	}

	return bytes
}
