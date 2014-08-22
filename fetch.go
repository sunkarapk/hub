package main

import (
	"strings"
)

type FetchCommand struct{}

func (f *FetchCommand) Execute(args []string) error {
	if len(args) == 0 {
		return &ErrArgument{}
	} else if len(args) > 1 {
		return &ErrProxy{}
	}

	if len(strings.Split(args[0], "/")) != 1 {
		return &ErrProxy{}
	}

	users := strings.Split(args[0], ",")
	remotes, err := Remotes()

	remoteAdd := &RemoteAddCommand{}

	if err != nil {
		return err
	}

	for _, user := range users {
		if _, ok := remotes[user]; !ok {
			remoteAdd.Execute([]string{user})
		}
	}

	err = Git([]string{"fetch", "--multiple", strings.Join(users, " ")}...)

	if err != nil {
		return err
	}

	HandleInfo("Fetched from remotes " + args[0])

	return nil
}

func (f *FetchCommand) Usage() string {
	return "<user | users>"
}
