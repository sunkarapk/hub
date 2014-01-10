package main

type AuthDestroyCommand struct{}

var AuthDestroy AuthDestroyCommand

func (a *AuthDestroyCommand) Execute(args []string) error {
	return nil
}

func (a *AuthDestroyCommand) Usage() string {
	return ""
}
