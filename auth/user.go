package auth

import (
	"github.com/pksunkara/hub/utils"
)

type UserCommand struct{}

func (a *UserCommand) Execute(args []string) error {
	if Config("token") != "" {
		return &utils.ErrModes{}
	}

	if len(args) == 1 {
		Config("user", args[0])
		utils.HandleInfo("You are now in `user` mode")
	} else {
		return &utils.ErrArgument{}
	}

	return nil
}

func (a *UserCommand) Usage() string {
	return "<user>"
}
