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
	"github.com/rjelierse/adventofcode-2017/06"
	"github.com/rjelierse/adventofcode-2017/07"
	"github.com/rjelierse/adventofcode-2017/08"
	"github.com/rjelierse/adventofcode-2017/09"
	"github.com/rjelierse/adventofcode-2017/10"
	"github.com/rjelierse/adventofcode-2017/11"
	"github.com/rjelierse/adventofcode-2017/12"
	"github.com/rjelierse/adventofcode-2017/13"
	"github.com/rjelierse/adventofcode-2017/14"
	"github.com/rjelierse/adventofcode-2017/15"
	"github.com/rjelierse/adventofcode-2017/16"
	"os"
	"github.com/rjelierse/adventofcode-2017/17"
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
	subcommands.Register(garbage.Command(), "Advent of Code 2017")
	subcommands.Register(knot.Command(), "Advent of Code 2017")
	subcommands.Register(hexes.Command(), "Advent of Code 2017")
	subcommands.Register(pipes.Command(), "Advent of Code 2017")
	subcommands.Register(firewall.Command(), "Advent of Code 2017")
	subcommands.Register(defrag.Command(), "Advent of Code 2017")
	subcommands.Register(generator.Command(), "Advent of Code 2017")
	subcommands.Register(dance.Command(), "Advent of Code 2017")
	subcommands.Register(spinlock.Command(), "Advent of Code 2017")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
