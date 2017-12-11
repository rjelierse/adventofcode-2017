package hexes

import (
	"github.com/google/subcommands"
	"flag"
	"context"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

type command struct {
	path string
}

func (*command) Name() string {
	return "day11"
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
	input, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read input", err)
	}
	moves := strings.Split(strings.TrimSpace(string(input)), ",")

	var dist, max int
	g := NewGrid()
	for _, move := range moves {
		g.Move(move)
		dist = g.Distance()

		if dist > max {
			max = dist
		}
	}

	fmt.Println("Largest distance:", max)
	fmt.Println("Final distance:", dist)

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
