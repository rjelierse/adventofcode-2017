package day20

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"os"
	"strings"
	"sort"
)

type command struct {
	path   string
	rounds int
}

func (c *command) Name() string {
	return "day20"
}

func (c *command) Synopsis() string {
	return ""
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "input", "", "")
	f.IntVar(&c.rounds, "rounds", 500, "")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read input:", err)
		return subcommands.ExitFailure
	}

	input := strings.Split(strings.TrimSpace(string(data)), "\n")
	if err := c.part1(input); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to run part 1:", err)
		return subcommands.ExitFailure
	}

	if err := c.part2(input); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to run part 2:", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func (c *command) part1(input []string) error {
	buffer, err := BufferFromInput(input)
	if err != nil {
		return err
	}

	fmt.Println("Closest point:", buffer.Closest(c.rounds).Id)
	return nil
}

func (c *command) part2(input []string) error {
	buffer, err := BufferFromInput(input)
	if err != nil {
		return err
	}

	sort.Slice(buffer.particles, func (i, j int) bool {
		a, b := buffer.particles[i], buffer.particles[j]

		switch {
		case a.Acceleration.Sum() < b.Acceleration.Sum():
			return true
		case a.Velocity.Sum() < b.Velocity.Sum():
			return true
		case a.Position.Sum() < b.Position.Sum():
			return true
		default:
			return false
		}
	})

	for t := 0; buffer.Run(t); t++ {
		if t % 1000 == 0 {
			fmt.Printf("[%05d] collisions: %d\n", t, buffer.Count(true))
		}
	}

	fmt.Println("Particles remaining:", buffer.Count(false))

	return nil
}

func Command() subcommands.Command {
	return &command{}
}
