package captcha

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
)

type command struct {
	file string
}

func (c *command) Name() string {
	return "day1"
}

func (c *command) Synopsis() string {
	return "Get the results for day 1"
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(flags *flag.FlagSet) {
	flags.StringVar(&c.file, "input", "", "The file to read the input from")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	input := readInput(c.file)

	// Part 1:
	fmt.Println("[1] Captcha result:", calculateSum(input, 1))

	// Part 2:
	fmt.Println("[2] Captcha result:", calculateSum(input, len(input)/2))

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
