package main

type RemoteAddCommand struct {
	Private bool `short:"p" long:"private" description:"Use private url for the repository"`
}

func (r *RemoteAddCommand) Execute(args []string) error {
	return nil
}

func (r *RemoteAddCommand) Usage() string {
	return "[-p] <user | origin>"
}
