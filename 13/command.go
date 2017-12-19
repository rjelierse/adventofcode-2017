package day13

import (
	"flag"
	"github.com/google/subcommands"
	"context"
	"io/ioutil"
	"fmt"
	"os"
	"bytes"
)

type command struct {
	path string
}

func (*command) Name() string {
	return "day13"
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
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read input:", err)
		return subcommands.ExitFailure
	}

	firewall := NewFirewall()
	for _, input := range bytes.Split(bytes.TrimSpace(data), []byte("\n")) {
		var d, r int
		_, err := fmt.Sscanf(string(input), "%d: %d", &d, &r)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot read input:", err)
			return subcommands.ExitFailure
		}

		firewall.AddLayer(d, r)
	}

	fmt.Println("Penalty when traversing:", firewall.Traverse(0))

	for delay := 0;; delay++ {
		penalty := firewall.Traverse(delay)

		if penalty == 0 {
			fmt.Printf("Can traverse without being caught at delay %d\n", delay)
		}
	}

	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
