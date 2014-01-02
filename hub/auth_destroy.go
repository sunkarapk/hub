package hub

type AuthDestroy struct {
	Common
}

func (a *AuthDestroy) Short() string {
	return "d"
}

func (a *AuthDestroy) Title() string {
	return "destroy"
}

func (a *AuthDestroy) Usage() string {
	return ""
}

func (a *AuthDestroy) Description() string {
	return "Destroy authorization and delete username"
}

func (a *AuthDestroy) Run(args []string) {
}

func AuthDestroyCommand() Command {
	cmd := new(AuthDestroy)

	return cmd
}
