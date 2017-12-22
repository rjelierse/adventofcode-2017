package day19

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"os"
)

type command struct {
	path string
}

func (c *command) Name() string {
	return "day19"
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
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read input:", err)
		return subcommands.ExitFailure
	}

	maze := NewMaze(bytes.Split(data, []byte("\n")))
	sequence, steps := maze.Travel()
	fmt.Println("Sequence:", string(sequence))
	fmt.Println("Steps:", steps)

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
