package main

type RemoteCommand struct {
	RemoteAdd RemoteAddCommand `command:"add" description:"Add a github remote easily"`
}

func (r *RemoteCommand) Execute(args []string) error {
	return nil
}

func (r *RemoteCommand) Usage() string {
	return ""
}
