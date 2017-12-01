package captcha

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"time"
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
	start := time.Now()

	input := readInput(c.file)
	atInput := time.Now()

	// Part 1:
	sum1 := calculateSum(input, 1)
	atPart1 := time.Now()

	// Part 2:
	sum2 := calculateSum(input, len(input)/2)
	atPart2 := time.Now()

	fmt.Printf("Captchas generated (input parsed in %dµs).\n", atInput.Sub(start).Nanoseconds() / 1000)
	fmt.Printf("Part 1: %d (generated in %dµs)\n", sum1, atPart1.Sub(atInput).Nanoseconds() / 1000)
	fmt.Printf("Part 2: %d (generated in %dµs)\n", sum2, atPart2.Sub(atPart1).Nanoseconds() / 1000)

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
