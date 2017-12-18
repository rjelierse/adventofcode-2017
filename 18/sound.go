package duet

import (
	"errors"
	"strconv"
	"strings"
)

func Sounds(instructions []string) (lastPlayed int, err error) {
	registers := make(map[string]int)

	get := func(s string) int {
		if strings.IndexAny(s, "0123456789") != -1 {
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return v
		}
		return registers[s]
	}

	var pos int
	for pos >= 0 && pos < len(instructions) {
		fields := strings.Fields(instructions[pos])
		switch fields[0] {
		case "jgz":
			if get(fields[1]) > 0 {
				pos += get(fields[2])
				continue
			}
		case "rcv":
			if get(fields[1]) != 0 {
				return
			}
		case "snd":
			lastPlayed = get(fields[1])
		case "set":
			registers[fields[1]] = get(fields[2])
		case "add":
			registers[fields[1]] += get(fields[2])
		case "mul":
			registers[fields[1]] *= get(fields[2])
		case "mod":
			registers[fields[1]] %= get(fields[2])
		default:
			err = errors.New("Unknown command:" + fields[0])
			return
		}
		pos++
	}
	return 0, errors.New("program terminated without recovery")
}
