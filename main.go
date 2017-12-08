package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode-2017/01"
	"github.com/rjelierse/adventofcode-2017/02"
	"github.com/rjelierse/adventofcode-2017/03"
	"github.com/rjelierse/adventofcode-2017/04"
	"github.com/rjelierse/adventofcode-2017/05"
	"os"
	"github.com/rjelierse/adventofcode-2017/06"
	"github.com/rjelierse/adventofcode-2017/07"
	"github.com/rjelierse/adventofcode-2017/08"
)

func main() {
	// Register default commands
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.HelpCommand(), "")

	subcommands.Register(captcha.Command(), "Advent of Code 2017")
	subcommands.Register(checksum.Command(), "Advent of Code 2017")
	subcommands.Register(spiral.Command(), "Advent of Code 2017")
	subcommands.Register(passphrase.Command(), "Advent of Code 2017")
	subcommands.Register(trampoline.Command(), "Advent of Code 2017")
	subcommands.Register(banks.Command(), "Advent of Code 2017")
	subcommands.Register(tower.Command(), "Advent of Code 2017")
	subcommands.Register(registers.Command(), "Advent of Code 2017")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
