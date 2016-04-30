package main

import (
	"github.com/pksunkara/hub/config"
	"github.com/pksunkara/hub/utils"
	"strings"
)

type CloneCommand struct {
	Private bool `short:"p" long:"private" description:"Use private url for the repository"`
}

func (c *CloneCommand) Execute(args []string) error {
	if len(args) == 0 {
		return &utils.ErrArgument{}
	}

	if strings.Index(args[0], ":") != -1 {
		return &utils.ErrProxy{}
	}

	var repo string

	if c.Private {
		repo = "git@" + config.Get("site") + ":"
	} else {
		repo = "git://" + config.Get("site") + "/"
	}

	path := strings.Split(args[0], "/")

	if len(path) == 1 {
		if config.Get("user") == "" {
			return &utils.ErrUserMode{}
		}

		if config.Get("token") != "" {
			repo = "git@" + config.Get("site") + ":"
		}

		repo = repo + config.Get("user") + "/" + path[0]
	} else if len(path) == 2 {
		repo = repo + args[0]
	} else {
		return &utils.ErrProxy{}
	}

	return utils.Git([]string{"clone", "--progress", repo}...)
}

func (c *CloneCommand) Usage() string {
	return "[-p] [<user>/]<repo>"
}
