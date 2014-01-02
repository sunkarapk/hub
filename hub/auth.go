package hub

type Auth struct {
	Common
}

func (a *Auth) Short() string {
	return "a"
}

func (a *Auth) Title() string {
	return "auth"
}

func (a *Auth) Usage() string {
	return "<command> [<args>]"
}

func (a *Auth) Description() string {
	return "Manage github access modes"
}

func (a *Auth) Run(args []string) {
	if len(args) == 0 {
		ShowHelp(a)
	} else {
		a.children[args[0]].Run(args[1:])
	}
}

func AuthCommand() Command {
	cmd := new(Auth)

	cmd.Init()

	cmd.Add(AuthDestroyCommand())
	cmd.Add(AuthPrivateCommand())
	cmd.Add(AuthUserCommand())

	return cmd
}
