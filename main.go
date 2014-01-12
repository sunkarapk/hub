package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

const (
	HubVersion = "0.1.0"
)

// Options for flags package
var Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	Auth    AuthCommand    `command:"auth" description:"Manage github access modes"`
	Clone   CloneCommand   `command:"clone" description:"Helps cloning github repos"`
	Version VersionCommand `command:"version" description:"Display program version"`
}

var parser = flags.NewParser(&Options, flags.HelpFlag|flags.PassDoubleDash)

func main() {
	// Set usage string
	parser.Usage = "[-v]"

	// Parse the arguments
	args, err := parser.Parse()

	if err != nil {
		if _, ok := err.(*flags.Error); !ok {
			fmt.Fprintln(os.Stderr, "Error:", err)
			fmt.Fprintln(os.Stderr)
		} else {
			typ := err.(*flags.Error).Type

			if typ == flags.ErrUnknownCommand {
				err = errors.New("unknown command '" + args[0] + "'")
			}

			if typ != flags.ErrCommandRequired && typ != flags.ErrHelp {
				fmt.Fprintln(os.Stderr, err)
			}
		}

		parser.WriteHelp(os.Stderr)
	}
}
