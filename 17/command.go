package spinlock

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
	ring.Run(c.steps, 50000000)
	fmt.Println("Value after 0, 50000000 rounds:", ring.ValueAfter(0))
	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}

