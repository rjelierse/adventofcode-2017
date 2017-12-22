package day22

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"os"
)

type command struct {
	path string
}

func (c *command) Name() string {
	return "day22"
}

func (c *command) Synopsis() string {
	return ""
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "input", "", "")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	input, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot parse input:", err)
		return subcommands.ExitFailure
	}

	fmt.Println("[part1] Infections:", c.part1(input, 1e4))
	fmt.Println("[part2] Infections:", c.part2(input, 1e7))

	return subcommands.ExitSuccess
}

func (c *command) part1(input []byte, rounds int) int {
	m := NewMap(input)
	v := NewVirus()

	for i := 0; i < rounds; i++ {
		switch m.GetState(v.Position) {
		case StateInfected:
			v.TurnRight()
			m.SetState(v.Position, StateClean)
		case StateClean:
			v.TurnLeft()
			m.SetState(v.Position, StateInfected)
			v.AddInfection()
		default:
			panic("unknown state")
		}
		v.Move()
	}
	return v.Infections
}

func (c *command) part2(input []byte, rounds int) int {
	m := NewMap(input)
	v := NewVirus()

	for i := 0; i < rounds; i++ {
		switch m.GetState(v.Position) {
		case StateInfected:
			v.TurnRight()
			m.SetState(v.Position, StateFlagged)
		case StateClean:
			v.TurnLeft()
			m.SetState(v.Position, StateWeakened)
		case StateWeakened:
			m.SetState(v.Position, StateInfected)
			v.AddInfection()
		case StateFlagged:
			v.Reverse()
			m.SetState(v.Position, StateClean)
		default:
			panic("unknown state")
		}
		v.Move()
	}
	return v.Infections
}

func Command() subcommands.Command {
	return &command{}
}
