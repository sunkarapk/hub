package config

type AddCommand struct{}

func (a *AddCommand) Execute(args []string) {
	if len(args) == 2 {
		config.Set(args[0], args[1])
		utils.HandleInfo("Configuration changed successfully")
	} else {
		return &utils.ErrArgument{}
	}

	return nil
}

func (a *AddCommand) Usage() string {
	return "<key> <value>"
}
