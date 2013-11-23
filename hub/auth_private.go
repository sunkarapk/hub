package hub

import (
	"github.com/howeyc/gopass"
)

type AuthPrivate struct {
	Common
}

func (a *AuthPrivate) Short() string {
	return "p"
}

func (a *AuthPrivate) Title() string {
	return "private"
}

func (a *AuthPrivate) Description() string {
	return "Give access to your private data"
}

func (a *AuthPrivate) Usage() string {
	return "<username>"
}

func (a *AuthPrivate) Run(args []string) {
}

func AuthPrivateCommand() Command {
	cmd := new(AuthPrivate)

	return cmd
}
