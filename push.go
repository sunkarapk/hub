package main

import (
	"github.com/pksunkara/hub/utils"
	"strings"
)

type PushCommand struct{}

func (p *PushCommand) Execute(args []string) error {
	if len(args) != 2 {
		return &utils.ErrProxy{}
	}

	remotes := strings.Split(args[0], ",")

	if len(remotes) == 1 {
		return &utils.ErrProxy{}
	}

	for _, remote := range remotes {
		err := utils.Git(append([]string{"push", remote}, args[1:]...)...)

		if err != nil {
			return err
		}

		utils.HandleInfo("Pushed to `" + remote + "` remote")
	}

	return nil
}
