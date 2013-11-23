package hub

import (
	"fmt"
	"reflect"
)

type Help struct {
	commands map[string]Command
	options  reflect.Type
}

func (h *Help) Short() string {
	return "h"
}

func (h *Help) Title() string {
	return "help"
}

func (h *Help) Description() string {
	return "Show help messages of commands"
}

func (h *Help) Usage() string {
	return "hub [<options>] <command> [<args>]"
}

func (h *Help) mainHelp() {
	help := "USAGE:\n hub [<options>] <command> [<args>]\n\n"
	help += "  Command line tool to interact with github\n\nCOMMANDS:\n"

	for key, val := range h.commands {
		if len(key) == 1 {
			continue
		}

		help += "  " + val.Short() + ", " + val.Title()
		//TODO: smart padding
		help += "\t" + val.Description() + "\n"
	}

	h.optionsHelp(&help)

	help += "\nRun 'hub help <command>' for more information on a specific comamnd."

	fmt.Println(help)
}

func (h *Help) optionsHelp(help *string) {
	*help += "\nOPTIONS:\n"

	for i := 0; i < h.options.NumField(); i++ {
		tag := h.options.Field(i).Tag
		*help += "  -" + tag.Get("short") + ", --" + tag.Get("long") + "\t  " + tag.Get("description") + "\n"
	}
}

func (h *Help) Run(args []string) {

	if len(args) == 0 || args[0] == "help" {
		h.mainHelp()
	} else {
		if _, err := h.commands[args[0]]; !err {
			fmt.Println("Command does not exist: '" + args[0] + "'\n")
			h.mainHelp()
		} else {
			fmt.Println("Work in Progress!")
		}
	}
}

func HelpCommand(cmds map[string]Command, opts reflect.Type) Command {
	cmd := new(Help)

	cmd.commands = cmds
	cmd.options = opts

	return cmd
}
