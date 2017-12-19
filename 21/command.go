package day21

import (
	"context"
	"flag"
	"github.com/google/subcommands"
)

type command struct{}

func (c *command) Name() string {
	return "day21"
}

func (c *command) Synopsis() string {
	return ""
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	panic("implement me")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	panic("implement me")
}

func Command() subcommands.Command {
	return &command{}
}
