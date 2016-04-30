package auth

import (
	"github.com/pksunkara/hub/config"
	"github.com/pksunkara/hub/utils"
)

type DestroyCommand struct{}

func (d *DestroyCommand) Execute(args []string) error {
	config.Set("user", "")
	config.Set("token", "")

	utils.HandleInfo("You are now in `public` mode")

	return nil
}
