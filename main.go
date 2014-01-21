package main

import (
	"errors"
	"github.com/jessevdk/go-flags"
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

	Auth    AuthCommand    `command:"auth" description:"Manage github access modes"`
	Clone   CloneCommand   `command:"clone" description:"Clone github repos easily"`
	Fetch   FetchCommand   `command:"fetch" description:"Fetch user's repo updates"`
	Fork    ForkCommand    `command:"fork" description:"Fork a github repo"`
	Remote  RemoteCommand  `command:"remote" description:"Manage remotes of repos"`
	Version VersionCommand `command:"version" description:"Display program version"`
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

		if len(option) > 0 {
			if len(option) > 1 {
				conf.AddOption("default", option[0], option[1])
				conf.WriteFile(hubrc, 0600, "Config for http://github.com/pksunkara/hub")
				conf, _ = config.ReadDefault(hubrc)

				value = option[1]
			} else {
				value, _ = conf.String("default", option[0])
			}
		}

		return value
	}

	// Parse the arguments
	args, err := parser.Parse()

	if err != nil {
		if Config("combine") == "1" {
			err := Git(os.Args[1:]...)

			if err != nil {
				os.Exit(1)
			} else {
				os.Exit(0)
			}
		}

		if _, ok := err.(*exec.ExitError); ok {
			HandleError(errors.New("Running git command is unsuccessful"))
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

		HandleError(err)

		terminal.Stderr.Nl()
		parser.WriteHelp(os.Stderr)
	}
}

func HandleError(err error) {
	if err != nil {
		terminal.Stderr.Color("r").Print("errs").Color("w!").Print(": ", err).Reset().Nl()
	}
}

func HandleInfo(str string) {
	terminal.Stderr.Color("g").Print("info").Color("w!").Print(": ", str).Reset().Nl()
}

func HandleDebug(str string) {
	if Options.Verbose {
		terminal.Stderr.Color("y").Print("logs").Color("w!").Print(": ", str).Reset().Nl()
	}
}

func Git(args ...string) error {
	if _, err := exec.LookPath("git"); err != nil {
		HandleError(errors.New("Please install git on your system"))
		os.Exit(1)
	}

	cmd := exec.Command("git", args...)

	cmd.Stdin = os.Stdin

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	HandleDebug("git " + strings.Join(args, " "))

	return cmd.Run()
}
