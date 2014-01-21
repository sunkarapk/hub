package main

type AuthUserCommand struct{}

func (a *AuthUserCommand) Execute(args []string) error {
	if Config("token") != "" {
		return &ErrModes{}
	}

	if len(args) == 1 {
		Config("user", args[0])
		HandleInfo("You are now in `user` mode")
	} else {
		return &ErrArgument{}
	}

	return nil
}

func (a *AuthUserCommand) Usage() string {
	return "<user>"
}
