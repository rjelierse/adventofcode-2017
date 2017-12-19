package day17

import (
	"flag"
	"github.com/google/subcommands"
	"context"
	"fmt"
)

type command struct {
	steps int
}

func (c *command) Name() string {
	return "day17"
}

func (c *command) Synopsis() string {
	return ""
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.IntVar(&c.steps, "steps", 3, "")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	ring := NewRing()
	ring.Run(c.steps, 2017)
	fmt.Println("Value after 2017, 2017 rounds:", ring.ValueAfter(2017))

	index := 0
	result := 0
	for value := 1; value < 50e6; value++ {
		index = (index + c.steps + 1) % value
		if index == 0 {
			result = value
		}
	}
	fmt.Println("Value after 0, 50000000 rounds:", result)
	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}

