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

	floor := NewDancefloor(16)

	for i := 0; i < c.rounds; i++ {
		floor.Dance(instructions)
		if iter := floor.IsRepeat(); iter > -1 && iter != i {
			round := (c.rounds % i) - 1
			fmt.Printf("Round %d is a repeat of round %d: %s\n", i+1, iter+1, floor.history[iter])
			fmt.Printf("Round %d is a repeat of round %d: %s\n", c.rounds, round+1, floor.history[round])
			return subcommands.ExitSuccess
		}
	}

	fmt.Printf("Dancefloor positions: %s\n", string(floor.Positions))

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
