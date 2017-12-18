package sound

import (
	"errors"
	"strconv"
)

func SoundBoard(instructions []string) (int, error) {
	var lastPlayed int
	registers := make(map[string]int)
	for i := 0; i >= 0 && i < len(instructions); {
		cmd := instructions[i][0:3]
		addr := instructions[i][4:5]

		if cmd == "rcv" && registers[addr] > 0 {
			return lastPlayed, nil
		}
		if cmd == "rcv" && registers[addr] <= 0 {
			i++
			continue
		}

		if cmd == "jgz" && registers[addr] != 0 {
			value, err := strconv.Atoi(instructions[i][6:])
			if err != nil {
				return 0, err
			}
			i = i + value
			continue
		}
		if cmd == "jgz" && registers[addr] == 0 {
			i++
			continue
		}

		if cmd == "snd" {
			lastPlayed = registers[addr]
			i++
			continue
		}

		value, err := strconv.Atoi(instructions[i][6:])
		if err != nil {
			addr := instructions[i][6:7]
			value = registers[addr]
		}

		switch cmd {
		case "set":
			registers[addr] = value
		case "add":
			registers[addr] = registers[addr] + value
		case "mul":
			registers[addr] = registers[addr] * value
		case "mod":
			registers[addr] = registers[addr] % value
		}

		i++
	}
	return 0, errors.New("program terminated without recovery")
}
