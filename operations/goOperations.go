package operations

import (
	"os/exec"
)

func RunTidy() (string, error) {
	err := ReadSaveDeleteDefaultServers_WriteNewServers()
	if err != nil {
		return "", err
	}

	cmd := exec.Command("go", "mod", "tidy")

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	err = RestoreResolv()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
