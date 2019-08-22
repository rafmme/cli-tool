package runcmd

import (
	"os/exec"
	"strings"
)

func Execute(path, filename string) (*string,error) {
	out, err := exec.Command(
		"bash", "-c",
		"cd "+ path + " && touch " + filename +
		"; pwd",
	).Output()
	fullPath := strings.TrimSpace(string(out[:])) + "/" + filename

	if err != nil {
		return nil, err
	}
	return &fullPath, nil
}
