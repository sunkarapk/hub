package main

import (
	"errors"
	"github.com/jessevdk/go-flags"
	"github.com/pksunkara/hub/auth"
	"github.com/pksunkara/hub/config"
	"github.com/pksunkara/hub/generate"
	"github.com/pksunkara/hub/remote"
	"github.com/pksunkara/hub/utils"
	"github.com/wsxiaoys/terminal"
	"os"
	"os/exec"
	"strconv"
)

const (
	HubVersion = "0.1.0"
)

// Options for flags package
var Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	Auth     auth.Command     `command:"auth" alias:"a" description:"Manage github access modes"`
	Clone    CloneCommand     `command:"clone" alias:"c" description:"Clone github repos easily"`
	Config   config.Command   `command:"config" description:"Manage application configuration"`
	Fetch    FetchCommand     `command:"fetch" description:"Fetch multiple users repo updates"`
	Fork     ForkCommand      `command:"fork" description:"Fork a github repo"`
	Generate generate.Command `command:"generate" alias:"g" description:"Generate snippets in repos"`
	Push     PushCommand      `command:"push" description:"Push to multiple github repos"`
	Remote   remote.Command   `command:"remote" alias:"r" description:"Manage remotes of repos" subcommands-optional:"1"`
	Version  VersionCommand   `command:"version" description:"Display application version"`
}

func main() {
	var err error

	// Initiate parser
	parser := flags.NewParser(&Options, flags.HelpFlag|flags.PassDoubleDash)

	// Set usage string
	parser.Usage = "[-v]"

	// Load config for application
	config.ItExists()

	// Parse the arguments
	args, err := parser.Parse()

	// Check for verbose mode
	os.Setenv("HUB_VERBOSE", strconv.FormatBool(Options.Verbose))

	if err != nil {
		if config.Get("combine") == "1" {
			err := utils.Git(os.Args[1:]...)

			if err != nil {
				os.Exit(1)
			} else {
				os.Exit(0)
			}
		}

		if _, ok := err.(*exec.ExitError); ok {
			utils.HandleError(errors.New("Running git command is unsuccessful"))
			os.Exit(1)
		}

		if _, ok := err.(*flags.Error); ok {
			typ := err.(*flags.Error).Type

			if typ == flags.ErrUnknownCommand {
				err = errors.New("unknown command '" + args[0] + "'")
			}

			if typ == flags.ErrCommandRequired || typ == flags.ErrHelp {
				err = nil
			}
		}

		utils.HandleError(err)

		terminal.Stderr.Nl()
		parser.WriteHelp(os.Stderr)
	}
}
