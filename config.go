package main

type ConfigCommand struct {
	ConfigAdd ConfigAddCommand `command:"add" alias:"a" description:"Add a new variable"`
}
