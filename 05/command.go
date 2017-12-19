package day05

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"strconv"
	"strings"
)

type command struct {
	path  string
	moves []int
}

func (c *command) Name() string {
	return "day5"
}

func (c *command) Synopsis() string {
	return "Solve assignment for day 5"
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "input", "", "Path to puzzle input")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if err := c.parseInput(); err != nil {
		return subcommands.ExitFailure
	}

	var steps int

	steps = CalculateSteps(c.moves)
	fmt.Println("Steps:", steps)

	steps = CalculateSteps2(c.moves)
	fmt.Println("Steps:", steps)

	return subcommands.ExitSuccess
}

func (c *command) parseInput() error {
	input, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	c.moves = make([]int, len(lines))

	for i, data := range lines {
		move, err := strconv.ParseInt(data, 10, 0)
		if err != nil {
			return err
		}

		c.moves[i] = int(move)
	}

	return nil
}

func Command() subcommands.Command {
	return &command{}
}
