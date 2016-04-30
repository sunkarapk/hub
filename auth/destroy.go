package auth

import (
	"github.com/pksunkara/hub/utils"
)

type DestroyCommand struct{}

func (a *DestroyCommand) Execute(args []string) error {
	Config("user", "")
	Config("token", "")

	utils.HandleInfo("You are now in `public` mode")

	return nil
}
