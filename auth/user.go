package auth

import (
	"github.com/pksunkara/hub/config"
	"github.com/pksunkara/hub/utils"
)

type UserCommand struct{}

func (u *UserCommand) Execute(args []string) error {
	if config.Get("token") != "" {
		return &utils.ErrModes{}
	}

	if len(args) == 1 {
		config.Set("user", args[0])
		utils.HandleInfo("You are now in `user` mode")
	} else {
		return &utils.ErrArgument{}
	}

	return nil
}

func (u *UserCommand) Usage() string {
	return "<user>"
}
