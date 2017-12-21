package day20

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

	buffer, err := BufferFromInput(strings.Split(strings.TrimSpace(string(data)), "\n"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse input:", err)
		return subcommands.ExitFailure
	}

	fmt.Println("Closest point:", buffer.Closest(c.rounds).Id)

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
