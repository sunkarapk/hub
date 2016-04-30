package config

import (
	"github.com/robfig/config"
	"os/user"
	"path/filepath"
)

var (
	conf  *config.Config
	hubrc string
)

type Command struct {
	Add AddCommand `command:"add" alias:"a" description:"Add a new variable"`
}

func ItExists() {
	usr, _ := user.Current()
	hubrc = filepath.Join(usr.HomeDir, ".hubrc")

	var err error

	// Load config for application
	conf, err = config.ReadDefault(hubrc)

	if err != nil {
		conf = config.NewDefault()

		conf.AddSection("default")
		conf.AddOption("default", "site", "github.com")
		conf.AddOption("default", "combine", "0")

		conf.WriteFile(hubrc, 0600, "Config for http://github.com/pksunkara/hub")
		conf, _ = config.ReadDefault(hubrc)
	}
}

func Get(key string) string {
	ItExists()

	value, _ := conf.String("default", key)

	return value
}

func Set(key string, value string) string {
	ItExists()

	conf.AddOption("default", key, value)
	conf.WriteFile(hubrc, 0600, "Config for http://github.com/pksunkara/hub")

	conf, _ = config.ReadDefault(hubrc)

	return value
}
