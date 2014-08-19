package main

type ErrProxy struct{}

func (e *ErrProxy) Error() string {
	return "Set `combine: 1` in hubrc to use git when hub fails!"
}

type ErrUserMode struct{}

func (e *ErrUserMode) Error() string {
	return "You need to be in 'user' mode or 'private' mode!"
}

type ErrPrivateMode struct{}

func (e *ErrPrivateMode) Error() string {
	return "You need to be in 'private' mode!"
}

type ErrArgument struct{}

func (e *ErrArgument) Error() string {
	return "You are missing an expected argument or giving excessive arguments!"
}

type ErrModes struct{}

func (e *ErrModes) Error() string {
	return "Please destroy the current user before changing modes!"
}
