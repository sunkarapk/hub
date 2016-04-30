package auth

// import (
// 	"fmt"
// 	"github.com/howeyc/gopass"
// )

type PrivateCommand struct{}

func (a *PrivateCommand) Execute(args []string) error {
	return nil
}

func (a *PrivateCommand) Usage() string {
	return "<user>"
}
