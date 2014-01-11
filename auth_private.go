package main

// import (
// 	"fmt"
// 	"github.com/howeyc/gopass"
// )

type AuthPrivateCommand struct{}

func (a *AuthPrivateCommand) Execute(args []string) error {
	return nil
}

func (a *AuthPrivateCommand) Usage() string {
	return ""
}
