package sound

import (
	"flag"
	"github.com/google/subcommands"
	"context"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

type command struct {
	path string
}

func (c *command) Name() string {
	return "day18"
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
		fmt.Fprintln(os.Stderr, "Cannot read input:", err)
		return subcommands.ExitFailure
	}
	instructions := strings.Split(strings.TrimSpace(string(input)), "\n")
	// Part 1
	fmt.Println(SoundBoard(instructions))
	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}


