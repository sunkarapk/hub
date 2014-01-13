package main

type ForkCommand struct{}

func (f *ForkCommand) Execute(args []string) error {
	return nil
}

func (f *ForkCommand) Usage() string {
	return "[repo]"
}
