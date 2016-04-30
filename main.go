package main

import (
	"errors"
	"github.com/jessevdk/go-flags"
	"github.com/pksunkara/hub/auth"
	"github.com/pksunkara/hub/utils"
	"github.com/robfig/config"
	"github.com/wsxiaoys/terminal"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	HubVersion = "0.1.0"
)

var (
	Config func(...string) string
)

// Options for flags package
var Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	Auth     auth.Command    `command:"auth" alias:"a" description:"Manage github access modes"`
	Clone    CloneCommand    `command:"clone" alias:"c" description:"Clone github repos easily"`
	Config   ConfigCommand   `command:"config" description:"Manage hub's configuration"`
	Fetch    FetchCommand    `command:"fetch" description:"Fetch multiple users repo updates"`
	Fork     ForkCommand     `command:"fork" description:"Fork a github repo"`
	Generate GenerateCommand `command:"generate" alias:"g" description:"Generate something in the repo"`
	Push     PushCommand     `command:"push" description:"Push to multiple github repos"`
	Remote   RemoteCommand   `command:"remote" alias:"r" description:"Manage remotes of repos" subcommands-optional:"1"`
	Version  VersionCommand  `command:"version" description:"Display program version"`
}

func main() {
	var err error

	// Initiate parser
	parser := flags.NewParser(&Options, flags.HelpFlag|flags.PassDoubleDash)

	// Set usage string
	parser.Usage = "[-v]"

	var conf *config.Config

	usr, _ := user.Current()
	hubrc := filepath.Join(usr.HomeDir, ".hubrc")

	// Load config for application
	conf, err = config.ReadDefault(hubrc)

	if err != nil {
		conf = config.NewDefault()

		conf.AddSection("default")
		conf.AddOption("default", "site", "github.com")
		conf.AddOption("default", "combine", "0")

		conf.WriteFile(hubrc, 0600, "Config for http://github.com/pksunkara/hub")
		conf, _ = config.ReadDefault(hubrc)
	}

	Config = func(option ...string) string {
		var value string

		if len(option) > 1 {
			conf.AddOption("default", option[0], option[1])
			conf.WriteFile(hubrc, 0600, "Config for http://github.com/pksunkara/hub")
			conf, _ = config.ReadDefault(hubrc)

			value = option[1]
		} else if len(option) > 0 {
			value, _ = conf.String("default", option[0])
		}

		return value
	}

	// Parse the arguments
	args, err := parser.Parse()

	if err != nil {
		if Config("combine") == "1" {
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
