package utils

import (
	"github.com/wsxiaoys/terminal"
)

func HandleError(err error) {
	if err != nil {
		terminal.Stderr.Color("r").Print("errs").Color("w!").Print(": ", err).Reset().Nl()
	}
}

func HandleInfo(str string) {
	terminal.Stderr.Color("g").Print("info").Color("w!").Print(": ", str).Reset().Nl()
}

func HandleDebug(str string, verbose bool) {
	if verbose {
		terminal.Stderr.Color("y").Print("logs").Color("w!").Print(": ", str).Reset().Nl()
	}
}
