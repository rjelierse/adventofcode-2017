package day02

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"strings"
)

type command struct {
	file string
}

func (c *command) Name() string {
	return "day2"
}

func (c *command) Synopsis() string {
	return "Get the results for day 2"
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(flags *flag.FlagSet) {
	flags.StringVar(&c.file, "input", "", "The file to read the input from")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	input := strings.Split(strings.TrimSpace(readInput(c.file)), "\n")
	var sum1, sum2 int

	for _, line := range input {
		numbers := convertLine(line)

		sum1 = sum1 + Difference(numbers)
		sum2 = sum2 + Division(numbers)
	}

	fmt.Printf("Checksum: %d\n", sum1)
	fmt.Printf("Checksum: %d\n", sum2)

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
