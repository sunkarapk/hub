package main

import (
	"errors"
)

type CloneCommand struct {
	Private bool `short:"p" long:"private" description:"Use private url for the repository"`
}

func (c *CloneCommand) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("Please specify a repo!")
	}

	return nil
}

func (c *CloneCommand) Usage() string {
	return "[-p] <repo>"
}
