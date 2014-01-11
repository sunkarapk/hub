package main

type AuthDestroyCommand struct{}

func (a *AuthDestroyCommand) Execute(args []string) error {
	return nil
}

func (a *AuthDestroyCommand) Usage() string {
	return ""
}
