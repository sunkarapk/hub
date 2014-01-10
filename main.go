package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

const (
	HubVersion = "0.1.0"
)

// Options for flags package
var opts struct {
	Private bool `short:"p" long:"private" description:"Use private url for the repository"`

	Verbose []bool `short:"V" long:"verbose" description:"Show verbose debug information"`

	Auth    AuthCommand    `command:"auth" description:"Manage github access modes"`
	Version VersionCommand `command:"version" description:"Display program version"`
}

var parser = flags.NewParser(&opts, flags.Default)

func main() {
	// Set usage string
	parser.Usage = "[options]"

	// Parse the arguments
	_, err := parser.Parse()

	if err != nil {
		if _, ok := err.(*flags.Error); !ok || err.(*flags.Error).Type != flags.ErrHelp {
			fmt.Println()
			parser.WriteHelp(os.Stdout)
		}
	}
}
