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
	}

	instructions := SplitInstructions(string(data))
	floor := NewDancefloor(16)

	for i := 0; i < c.rounds; i++ {
		floor.Dance(instructions)
	}

	fmt.Println("Dancefloor positions:", string(floor.Positions))

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
