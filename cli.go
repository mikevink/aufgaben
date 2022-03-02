package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
)

type CLIAddOptions struct {
	Tags     []string `short:"t" long:"tag" description:"Tag(s) for this task"`
	Project  string   `long:"project" description:"Bit like a tag, but only 1"`
	Priority string   `short:"p" long:"priority" description:"Yawn or hair-on-fire" choice:"zzz" choice:"yawn" choice:"hmm" choice:"sweat" choice:"panic"`
}

type CLIUpdateOptions struct {
	Hash string `short:"h" long:"hash" description:"Hash of the tag to update"`
}

type CLIOptions struct {
	Config string           `long:"config" description:"Path to configuration directory" env:"AUFGABEN_CONFIG_DIR"`
	Add    CLIAddOptions    `command:"add" alias:"a" description:"Add a new task"`
	Update CLIUpdateOptions `command:"update" alias:"u" description:"Update an existing task"`
}

type CLI struct {
	Options CLIOptions
}

func ParseCLI() CLI {
	ucd, _ := os.UserConfigDir()
	cd := filepath.Join(ucd, "aufgaben")
	cli := CLI{
		Options: CLIOptions{Config: cd},
	}
	p := flags.NewParser(&cli, flags.Default)
	_, err := p.Parse()
	if nil != err && !flags.WroteHelp(err) {
		fmt.Println("boom")
		fmt.Println(err)
		os.Exit(1)
	}
	return cli
}
