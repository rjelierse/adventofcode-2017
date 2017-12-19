package day15

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"fmt"
)

type command struct {
	valueA int
	valueB int
	rounds int
	shouldDivide bool
}

func (*command) Name() string {
	return "day15"
}

func (*command) Synopsis() string {
	return ""
}

func (*command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.IntVar(&c.valueA, "a", startA, "")
	f.IntVar(&c.valueB, "b", startB, "")
	f.IntVar(&c.rounds, "rounds", sampleSize, "")
	f.BoolVar(&c.shouldDivide, "division", false, "")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	judge := NewJudge(c.valueA, c.valueB)

	fmt.Println(judge.FindMatches(c.rounds, c.shouldDivide))

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}