package day07

import (
	"flag"
	"github.com/google/subcommands"
	"context"
	"io/ioutil"
	"strings"
	"fmt"
)

type command struct {
	path string
}

func (c *command) Name() string {
	return "day7"
}

func (c *command) Synopsis() string {
	return ""
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "input", "", "")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	lines, err := c.parseInput()
	if err != nil {
		return subcommands.ExitFailure
	}

	node := BuildTree(lines)
	fmt.Printf("Root: %s\n", node.Name)

	_, err = node.Weight()
	if err != nil {
		fmt.Printf("Imbalanced node: %v\n", err)
	}

	return subcommands.ExitSuccess
}

func (c *command) parseInput() (lines []string, err error) {
	bytes, err := ioutil.ReadFile(c.path)
	if err != nil {
		return
	}

	lines = strings.Split(string(bytes), "\n")

	return
}

func Command() subcommands.Command {
	return &command{}
}