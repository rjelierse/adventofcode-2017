package garbage

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"io/ioutil"
	"fmt"
)

type command struct {
	path string
}

func (*command) Name() string {
	return "day9"
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
		return subcommands.ExitFailure
	}

	fmt.Println(ReadStream(string(input)))

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
