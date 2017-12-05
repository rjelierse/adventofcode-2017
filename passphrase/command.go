package passphrase

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"os"
	"strings"
)

type command struct {
	path string
}

func Command() subcommands.Command {
	return &command{}
}

func (c *command) Name() string {
	return "day4"
}

func (c *command) Synopsis() string {
	return "Complete assignment for day 4"
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "input", "", "The puzzle input")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	input, err := c.readInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read input:", err)
		return subcommands.ExitFailure
	}

	phrases := strings.Split(input, "\n")

	fmt.Fprintln(os.Stdin, "Number of valid passphrases:", partOne(phrases))
	fmt.Fprintln(os.Stdin, "Number of valid passphrases:", partTwo(phrases))

	return subcommands.ExitSuccess
}

func (c *command) readInput() (string, error) {
	bytes, err := ioutil.ReadFile(c.path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func partOne(phrases []string) (valid int) {
	for _, phrase := range phrases {
		if IsUnique(phrase) {
			valid++
		}
	}

	return valid
}

func partTwo(phrases []string) (valid int) {
	for _, phrase := range phrases {
		if IsNotAnagram(phrase) {
			valid++
		}
	}

	return valid
}
