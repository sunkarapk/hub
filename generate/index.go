package generate

type Command struct {
	Readme ReadmeCommand `command:"readme" alias:"r" description:"Append some standard stuff to README.md"`
}
