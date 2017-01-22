package twbot

import (
	"os"

	"github.com/mitchellh/cli"
)

// UI is
var UI cli.Ui

const (
	infoPrefix  = "INFO: "
	warnPrefix  = "WARN: "
	errorPrefix = "ERROR: "
)

func init() {
	UI = &cli.PrefixedUi{
		InfoPrefix:  infoPrefix,
		WarnPrefix:  warnPrefix,
		ErrorPrefix: errorPrefix,
		Ui: &cli.BasicUi{
			Writer: os.Stdout,
		},
	}
}

// SearchCommandFactory is
func SearchCommandFactory() (cli.Command, error) {
	return &searchCommand{
		ui: UI,
	}, nil
}

// ReplyCommandFactory is
func ReplyCommandFactory() (cli.Command, error) {
	return &replyCommand{
		ui: UI,
	}, nil
}
