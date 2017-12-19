package day16

import (
	"strings"
)

func ParseInstructions(input string) (instructions []Instruction, err error) {
	inputs := splitInstructions(input)
	instructions = make([]Instruction, len(inputs))

	for i, instruction := range inputs {
		switch instruction[0] {
		case 's':
			instructions[i], err = ParseSpinInstruction(instruction)
		case 'x':
			instructions[i], err = ParseExchangeInstruction(instruction)
		case 'p':
			instructions[i] = ParsePartnerInstruction(instruction)
		}
		if err != nil {
			return nil, err
		}
	}

	return instructions, nil
}

func splitInstructions(input string) []string {
	return strings.Split(strings.TrimSpace(input), ",")
}
