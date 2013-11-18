package main

import (
	"./hub"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"reflect"
)

func main() {

	// Declaration of variables
	commands := make(map[string]hub.Command)

	// Options for flags package
	var opts struct {
		Help    bool   `short:"h" long:"help" description:"Show this help message"`
		Private bool   `short:"p" long:"private" description:"Use private url for the repository"`
		Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
		Version bool   `short:"V" long:"version" description:"Show version information"`
	}

	// Options struct metadata
	optsMeta := reflect.ValueOf(&opts).Type().Elem()

	// Parse the arguments
	args, err := flags.Parse(&opts)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Set help as default command
	if len(args) == 0 {
		args = []string{"help"}
	}

	// Fill all the commands
	commands["auth"] = hub.AuthCommand()

	// Fill the help command finally
	commands["help"] = hub.HelpCommand(commands, optsMeta)

	// Check if command exists
	if _, err := commands[args[0]]; !err {
		fmt.Println("Command does not exist: '" + args[0] + "'\n")
		args = []string{"help"}
	}

	// Run the command with given args
	commands[args[0]].Run(args[1:])
}
