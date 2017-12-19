package day07

import (
	"strings"
	"strconv"
)

func ParseLine(line string) (program *Program, err error) {
	parts := strings.Split(line, " ")

	weight, err := parseWeight(parts[1])
	if err != nil {
		return nil, err
	}

	program = new(Program)
	program.Name = parts[0]
	program.SelfWeight = weight
	program.Children = make([]string, 0)

	if len(parts) == 2 {
		return program, nil
	}

	for _, child := range parts[3:] {
		child = strings.TrimRight(child, ",")
		program.Children = append(program.Children, child)
	}

	return program, nil
}

func parseWeight(input string) (weight int, err error) {
	input = strings.Trim(input, "()")
	output, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		return
	}

	weight = int(output)

	return weight, nil
}
