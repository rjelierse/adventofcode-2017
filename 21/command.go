package day21

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

type command struct{
	path string
	rounds int
}

func (c *command) Name() string {
	return "day21"
}

func (c *command) Synopsis() string {
	return ""
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "input", "", "")
	f.IntVar(&c.rounds, "rounds", 2, "")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read input:", err)
		return subcommands.ExitFailure
	}
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	enh := ParseEnhancements(input)
	img := DefaultImage()
	for i := 0; i < c.rounds; i++ {
		img = img.Enhance(enh)
	}
	fmt.Println("Pixels on:", img.Count())

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
