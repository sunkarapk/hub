package utils

import (
	"github.com/wsxiaoys/terminal"
)

var Verbose *bool

func HandleError(err error) {
	if err != nil {
		terminal.Stderr.Color("r").Print("errs").Color("w!").Print(": ", err).Reset().Nl()
	}
}

func HandleInfo(str string) {
	terminal.Stderr.Color("g").Print("info").Color("w!").Print(": ", str).Reset().Nl()
}

func HandleDebug(str string) {
	if *Verbose {
		terminal.Stderr.Color("y").Print("logs").Color("w!").Print(": ", str).Reset().Nl()
	}
}
