package day18

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func program(pid int, instructions []string, out, in chan int) (sent int, err error) {
	registers := map[string]int{
		"p": pid,
	}

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
			select {
			case val := <-in:
				registers[fields[1]] = val
			case <-time.After(500 * time.Millisecond):
				return
			}
		case "snd":
			select {
			case out <- get(fields[1]):
				sent++
			case <-time.After(500 * time.Millisecond):
				return
			}
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

	return
}

func Run(instructions []string) (int, error) {
	c01 := make(chan int, 10000)
	c10 := make(chan int, 10000)
	go program(0, instructions, c01, c10)
	return program(1, instructions, c10, c01)
}
