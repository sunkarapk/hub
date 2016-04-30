package main

import (
	"github.com/pksunkara/hub/utils"
	"strings"
)

type RemoteAddCommand struct {
	Private bool `short:"p" long:"private" description:"Use private url for the repository"`
}

func (r *RemoteAddCommand) Execute(args []string) error {
	if len(args) == 0 {
		return &utils.ErrArgument{}
	} else if len(args) > 1 {
		return &utils.ErrProxy{}
	}

	var user, repo string

	if r.Private || args[0] == "origin" {
		repo = "git@" + Config("site") + ":"
	} else {
		repo = "git://" + Config("site") + "/"
	}

	if args[0] == "origin" {
		if Config("user") == "" {
			return &utils.ErrUserMode{}
		}

		user = Config("user")
	} else {
		user = args[0]
	}

	path := strings.Split(args[0], "/")

	if len(path) == 1 {
		repo = repo + user + "/" + Repo()
	} else {
		return &utils.ErrProxy{}
	}

	err := utils.Git([]string{"remote", "add", args[0], repo}...)

	if err != nil {
		return err
	}

	HandleInfo("Added remote named `" + args[0] + "`")

	return nil
}

func (r *RemoteAddCommand) Usage() string {
	return "[-p] <user | origin>"
}
