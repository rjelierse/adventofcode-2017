package knot

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"os"
	"bytes"
)

type command struct {
	path string
}

func (*command) Name() string {
	return "day10"
}

func (*command) Synopsis() string {
	return ""
}

func (*command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "input", "", "")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	lengths, err := c.getLengths()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot parse lengths", err)
		return subcommands.ExitFailure
	}

	for _, v := range Hash(lengths) {
		fmt.Fprintf(os.Stdout, "%02x", v)
	}

	fmt.Println()

	return subcommands.ExitSuccess
}

func (c *command) getLengths() ([]int, error) {
	input, err := ioutil.ReadFile(c.path)
	if err != nil {
		return nil, err
	}

	input = bytes.TrimSpace(input)

	lengths := make([]int, len(input))
	for i, value := range input {
		lengths[i] = int(value)
	}

	return lengths, nil
}

func Command() subcommands.Command {
	return &command{}
}
