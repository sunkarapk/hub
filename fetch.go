package main

import (
	"errors"
)

type FetchCommand struct {
	Private bool `short:"p" long:"private" description:"Use private url for the repository"`
}

func (f *FetchCommand) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("Please specify user(s)!")
	}

	return nil
}

func (f *FetchCommand) Usage() string {
	return "[-p] <user | users>"
}
