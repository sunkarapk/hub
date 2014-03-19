package main

import (
	"strings"
)

type CloneCommand struct {
	Private bool `short:"p" long:"private" description:"Use private url for the repository"`
}

func (c *CloneCommand) Execute(args []string) error {
	if len(args) == 0 {
		return &ErrArgument{}
	}

	if strings.Index(args[0], ":") != -1 {
		return &ErrProxy{}
	}

	var repo string

	if c.Private || Config("token") != "" {
		repo = "git@" + Config("site") + ":"
	} else {
		repo = "git://" + Config("site") + "/"
	}

	path := strings.Split(args[0], "/")

	if len(path) == 1 {
		if Config("user") == "" {
			return &ErrUserMode{}
		}

		repo = repo + Config("user") + "/" + path[0]
	} else if len(path) == 2 {
		repo = repo + args[0]
	} else {
		return &ErrProxy{}
	}

	return Git([]string{"clone", "--progress", repo}...)
}

func (c *CloneCommand) Usage() string {
	return "[-p] [<user>/]<repo>"
}
