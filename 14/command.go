package day14

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
)

type command struct {
	key string
}

func (*command) Name() string {
	return "day14"
}

func (*command) Synopsis() string {
	return ""
}

func (*command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.key, "key", "flqrgnkx", "")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	g := NewGrid(c.key)
	fmt.Println("Cells in use", g.Count())
	fmt.Println("Clusters", g.GroupCount())
	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
