package main

type GenerateCommand struct {
	GenerateReadme GenerateReadmeCommand `command:"readme" alias:"r" description:"Append some standard stuff to README.md"`
}
