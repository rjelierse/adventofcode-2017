package dance

import (
	"strings"
	"fmt"
	"log"
)

const (
	InstructionSpin = 's'
	InstructionExchange = 'x'
	InstructionPartner = 'p'
)

func SplitInstructions(input string) []string {
	return strings.Split(strings.TrimSpace(input), ",")
}

func GetInstruction(input string) byte {
	return input[0]
}

func ParseSpinInstruction(instruction string) (count int) {
	if _, err := fmt.Sscanf(instruction, "s%d", &count); err != nil {
		log.Fatal(err)
	}
	return
}

func ParseExchangeInstruction(instruction string) (posA int, posB int) {
	if _, err := fmt.Sscanf(instruction, "x%d/%d", &posA, &posB); err != nil {
		log.Fatal(err)
	}
	return
}

func ParsePartnerInstruction(instruction string) (a byte, b byte) {
	a = instruction[1]
	b = instruction[3]
	return
}
