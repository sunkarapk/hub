package hub

import (
	"fmt"
)

type Auth struct {
	children map[string]Command
}

func (a *Auth) Short() string {
	return "a"
}

func (a *Auth) Title() string {
	return "auth"
}

func (a *Auth) Description() string {
	return "Manage github access tokens"
}

func (a *Auth) Run(args []string) {
	fmt.Println(args)
}

func AuthCommand() Command {
	cmd := new(Auth)

	cmd.children = make(map[string]Command)

	return cmd
}
