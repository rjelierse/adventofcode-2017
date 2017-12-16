package dance

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"os"
)

type command struct {
	path   string
	rounds int
}

func (*command) Name() string {
	return "day16"
}

func (*command) Synopsis() string {
	return ""
}

func (*command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "input", "", "")
	f.IntVar(&c.rounds, "rounds", 1, "")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot parse input", err)
		return subcommands.ExitFailure
	}

	instructions, err := ParseInstructions(string(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot parse input", err)
		return subcommands.ExitFailure
	}

	floor := NewFloor(16)
	history := NewHistory()

	var position string
	for i := 0; i < c.rounds; i++ {
		position = floor.Dance(instructions)

		if seen, index := history.SeenBefore(position); seen {
			fmt.Printf("Sequence wraps at position %d\n", index)
			position = history.Get((c.rounds % i) - 1)
			break
		}

		history.Save(position)
	}

	fmt.Printf("Floor positions: %s\n", position)

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
