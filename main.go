package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode-2017/captcha"
	"github.com/rjelierse/adventofcode-2017/checksum"
	"os"
)

func main() {
	// Register default commands
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.HelpCommand(), "")

	subcommands.Register(captcha.Command(), "Advent of Code 2017")
	subcommands.Register(checksum.Command(), "Advent of Code 2017")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
