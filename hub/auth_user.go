package hub

type AuthUser struct {
	Common
}

func (a *AuthUser) Short() string {
	return "u"
}

func (a *AuthUser) Title() string {
	return "user"
}

func (a *AuthUser) Usage() string {
	return "<username>"
}

func (a *AuthUser) Description() string {
	return "Set an username to use their public data"
}

func (a *AuthUser) Run(args []string) {
}

func AuthUserCommand() Command {
	cmd := new(AuthUser)

	return cmd
}
