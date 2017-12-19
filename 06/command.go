package day06

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"strconv"
	"strings"
)

type command struct {
	path string
}

func (*command) Name() string {
	return "day6"
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
	config, err := c.readInput()
	if err != nil {
		return subcommands.ExitFailure
	}

	var step, first int
	var seen bool

	bank := NewBank(config)

	for step = 0; !seen; step++ {
		bank.Redistribute()
		seen, first = bank.ConfigurationSeenBefore()
	}

	fmt.Println("Steps", step, "length", step-first)

	return subcommands.ExitSuccess
}

func (c *command) readInput() (config Configuration, err error) {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		return config, err
	}

	slots := strings.Split(strings.TrimSpace(string(data)), "\t")
	config = make([]int, len(slots))
	for index, slot := range slots {
		value, err := strconv.ParseInt(slot, 10, 0)
		if err != nil {
			return config, err
		}
		config[index] = int(value)
	}

	return config, nil
}

func Command() subcommands.Command {
	return &command{}
}
