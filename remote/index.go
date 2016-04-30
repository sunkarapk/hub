package remote

import (
	"github.com/pksunkara/hub/utils"
)

type Command struct {
	Add AddCommand `command:"add" description:"Add a github remote easily"`
}

func (r *Command) Execute(args []string) error {
	return &utils.ErrProxy{}
}
