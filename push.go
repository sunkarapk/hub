package main

import (
	"strings"
)

type PushCommand struct{}

func (p *PushCommand) Execute(args []string) error {
	if len(args) != 2 {
		return &ErrProxy{}
	}

	remotes := strings.Split(args[0], ",")

	if len(remotes) == 1 {
		return &ErrProxy{}
	}

	for _, remote := range remotes {
		err := Git(append([]string{"push", remote}, args[1:]...)...)

		if err != nil {
			return err
		}

		HandleInfo("Pushed to `" + remote + "` remote")
	}

	return nil
}
