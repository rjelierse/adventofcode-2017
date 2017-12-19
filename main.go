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
	"github.com/rjelierse/adventofcode-2017/17"
	"github.com/rjelierse/adventofcode-2017/18"
	"github.com/rjelierse/adventofcode-2017/19"
	"os"
)

func main() {
	// Register default commands
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.HelpCommand(), "")

	subcommands.Register(day01.Command(), "Advent of Code 2017")
	subcommands.Register(day02.Command(), "Advent of Code 2017")
	subcommands.Register(day03.Command(), "Advent of Code 2017")
	subcommands.Register(day04.Command(), "Advent of Code 2017")
	subcommands.Register(day05.Command(), "Advent of Code 2017")
	subcommands.Register(day06.Command(), "Advent of Code 2017")
	subcommands.Register(day07.Command(), "Advent of Code 2017")
	subcommands.Register(day08.Command(), "Advent of Code 2017")
	subcommands.Register(day09.Command(), "Advent of Code 2017")
	subcommands.Register(day10.Command(), "Advent of Code 2017")
	subcommands.Register(day11.Command(), "Advent of Code 2017")
	subcommands.Register(day12.Command(), "Advent of Code 2017")
	subcommands.Register(day13.Command(), "Advent of Code 2017")
	subcommands.Register(day14.Command(), "Advent of Code 2017")
	subcommands.Register(day15.Command(), "Advent of Code 2017")
	subcommands.Register(day16.Command(), "Advent of Code 2017")
	subcommands.Register(day17.Command(), "Advent of Code 2017")
	subcommands.Register(day18.Command(), "Advent of Code 2017")
	subcommands.Register(day19.Command(), "Advent of Code 2017")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
