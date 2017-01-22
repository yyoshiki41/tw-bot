package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/yyoshiki41/tw-bot"
)

func main() {
	c := cli.NewCLI("twbot", twbot.Version())
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"search": twbot.SearchCommandFactory,
		"reply":  twbot.ReplyCommandFactory,
	}

	exitCode, err := c.Run()
	if err != nil {
		log.Printf("Error executing CLI: %s", err.Error())
	}

	os.Exit(exitCode)
}
