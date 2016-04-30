package main

import (
	"github.com/pksunkara/hub/utils"
	"strings"
)

type FetchCommand struct{}

func (f *FetchCommand) Execute(args []string) error {
	if len(args) == 0 {
		return &utils.ErrArgument{}
	} else if len(args) > 1 {
		return &utils.ErrProxy{}
	}

	if len(strings.Split(args[0], "/")) != 1 {
		return &utils.ErrProxy{}
	}

	users := strings.Split(args[0], ",")
	remotes, err := utils.Remotes()

	remoteAdd := &RemoteAddCommand{}

	if err != nil {
		return err
	}

	for _, user := range users {
		if _, ok := remotes[user]; !ok {
			remoteAdd.Execute([]string{user})
		}
	}

	err = utils.Git(append([]string{"fetch", "--multiple"}, users...)...)

	if err != nil {
		return err
	}

	utils.HandleInfo("Fetched from remotes " + args[0])

	return nil
}

func (f *FetchCommand) Usage() string {
	return "<user | users>"
}
