package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/yyoshiki41/tw-cli"
)

func main() {
	c := cli.NewCLI("twcli", twcli.Version())
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"reply": twcli.ReplyCommandFactory,
	}

	exitCode, err := c.Run()
	if err != nil {
		log.Printf("Error executing CLI: %s", err.Error())
	}

	os.Exit(exitCode)
}
