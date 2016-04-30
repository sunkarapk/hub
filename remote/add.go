package remote

import (
	"github.com/pksunkara/hub/config"
	"github.com/pksunkara/hub/utils"
	"strings"
)

type AddCommand struct {
	Private bool `short:"p" long:"private" description:"Use private url for the repository"`
}

func (a *AddCommand) Execute(args []string) error {
	if len(args) == 0 {
		return &utils.ErrArgument{}
	} else if len(args) > 1 {
		return &utils.ErrProxy{}
	}

	var user, repo string

	if a.Private || args[0] == "origin" {
		repo = "git@" + config.Get("site") + ":"
	} else {
		repo = "git://" + config.Get("site") + "/"
	}

	if args[0] == "origin" {
		if config.Get("user") == "" {
			return &utils.ErrUserMode{}
		}

		user = config.Get("user")
	} else {
		user = args[0]
	}

	path := strings.Split(args[0], "/")

	if len(path) == 1 {
		name, err := utils.RepoName()

		if err != nil {
			return err
		}

		repo = repo + user + "/" + name
	} else {
		return &utils.ErrProxy{}
	}

	err := utils.Git([]string{"remote", "add", args[0], repo}...)

	if err != nil {
		return err
	}

	utils.HandleInfo("Added remote named `" + args[0] + "`")

	return nil
}

func (a *AddCommand) Usage() string {
	return "[-p] <user | origin>"
}
