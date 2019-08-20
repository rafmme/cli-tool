package runcmd

import (
	"os/exec"
)

func Execute(path, filename string) ([]uint8,error) {
	out, err := exec.Command(
		"bash", "-c",
		"cd "+ path + " && touch " + filename +
		"; pwd",
	).Output()

	if err != nil {
		return nil, err
	}
	return out, nil
}
