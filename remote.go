package main

import (
	"github.com/pksunkara/hub/utils"
)

type RemoteCommand struct {
	RemoteAdd RemoteAddCommand `command:"add" description:"Add a github remote easily"`
}

func (r *RemoteCommand) Execute(args []string) error {
	return &utils.ErrProxy{}
}
