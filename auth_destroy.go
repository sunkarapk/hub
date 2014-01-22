package main

type AuthDestroyCommand struct{}

func (a *AuthDestroyCommand) Execute(args []string) error {
	Config("user", "")
	Config("token", "")

	HandleInfo("You are now in `public` mode")

	return nil
}
