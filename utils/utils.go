package utils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CheckGit() {
	if _, err := exec.LookPath("git"); err != nil {
		HandleError(errors.New("Please install git on your system"))
		os.Exit(1)
	}
}

func Repo() string {
	path, _ := os.Getwd()

	return filepath.Base(path)
}

func Remotes() (remotes map[string]string, err error) {
	CheckGit()

	remotes = make(map[string]string)

	out, cmderr := exec.Command("git", "remote", "-v").Output()

	if cmderr != nil {
		err = cmderr
		return
	}

	output := string(out[:])
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		if len(line) > 1 {
			remote := strings.Split(line, "\t")

			remotes[remote[0]] = strings.Split(remote[1], " ")[0]
		}
	}

	return
}

func Git(args ...string) error {
	CheckGit()

	cmd := exec.Command("git", args...)

	cmd.Stdin = os.Stdin

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	HandleDebug("git " + strings.Join(args, " "))

	return cmd.Run()
}
