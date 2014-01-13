package main

import (
	"errors"
	"github.com/jessevdk/go-flags"
	"github.com/wsxiaoys/terminal"
	"os"
)

const (
	HubVersion = "0.1.0"
)

// Options for flags package
var Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	Auth    AuthCommand    `command:"auth" description:"Manage github access modes"`
	Clone   CloneCommand   `command:"clone" description:"Clone github repos easily"`
	Fetch   FetchCommand   `command:"fetch" description:"Fetch user's repo updates"`
	Fork    ForkCommand    `command:"fork" description:"Fork a github repo"`
	Remote  RemoteCommand  `command:"remote" description:"Manage remotes of repos"`
	Version VersionCommand `command:"version" description:"Display program version"`
}

var parser = flags.NewParser(&Options, flags.HelpFlag|flags.PassDoubleDash)

func main() {
	// Set usage string
	parser.Usage = "[-v]"

	// Parse the arguments
	args, err := parser.Parse()

	if err != nil {
		if _, ok := err.(*flags.Error); ok {
			typ := err.(*flags.Error).Type

			if typ == flags.ErrUnknownCommand {
				err = errors.New("unknown command '" + args[0] + "'")
			}

			if typ == flags.ErrCommandRequired || typ == flags.ErrHelp {
				err = nil
			}
		}

		if err != nil {
			terminal.Stderr.Color("r!").Print("Error: ", err).Reset().Nl().Nl()
		}

		parser.WriteHelp(os.Stderr)
	}
}
