package day08

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"strings"
)

type Program struct {
	path  string
	lines []string
}

func (*Program) Name() string {
	return "day8"
}

func (*Program) Synopsis() string {
	return ""
}

func (*Program) Usage() string {
	return ""
}

func (p *Program) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.path, "input", "", "")
}

func (p *Program) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	input, err := ioutil.ReadFile(p.path)
	if err != nil {
		return subcommands.ExitFailure
	}

	p.lines = strings.Split(strings.TrimSpace(string(input)), "\n")
	maxValue, value := p.Run()

	fmt.Println(value, maxValue)

	return subcommands.ExitSuccess
}

func (p *Program) Run() (int, int) {
	var value int

	for _, line := range p.lines {
		instruction, condition := ParseLine(line)
		if condition.Valid() {
			instruction.Apply()
		}

		v := LargestValue()
		if v > value {
			value = v
		}
	}

	return value, LargestValue()
}

func Command() subcommands.Command {
	return &Program{}
}
