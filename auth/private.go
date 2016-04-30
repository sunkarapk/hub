package auth

// import (
// 	"fmt"
// 	"github.com/howeyc/gopass"
// )

type PrivateCommand struct{}

func (p *PrivateCommand) Execute(args []string) error {
	return nil
}

func (p *PrivateCommand) Usage() string {
	return "<user>"
}
