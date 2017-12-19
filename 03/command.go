package day03

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
)

type command struct {
	value int
}

func Command() subcommands.Command {
	return &command{}
}

func (c *command) Name() string {
	return "day3"
}

func (c *command) Synopsis() string {
	return "Complete assignment for day 3"
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.IntVar(&c.value, "input", 0, "The puzzle input")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	x, y := position(c.value)
	var carry int

	if x >= 0 && y >= 0 {
		carry = x + y
	} else if x >= 0 && y < 0 {
		carry = x + -1*y
	} else if x < 0 && y >= 0 {
		carry = -1*x + y
	} else if x < 0 && y < 0 {
		carry = -1*x + -1*y
	}

	fmt.Printf("Value at position (%d,%d) - carry %d\n", x, y, carry)

	return subcommands.ExitSuccess
}
