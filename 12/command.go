package day12

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"os"
	"strings"
)

type command struct {
	path string
}

func (c *command) Name() string {
	return "day12"
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
	input, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot parse input", err)
		return subcommands.ExitFailure
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	tree, err := ParseTree(lines)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot parse tree", err)
		return subcommands.ExitFailure
	}

	var pid int
	f.IntVar(&pid, "pid", 0, "")

	fmt.Println(tree.Size(pid), tree.NumGroups())

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
