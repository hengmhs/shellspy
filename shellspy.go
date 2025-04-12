package shellspy

import (
	"io"
	"os/exec"
)

func CommandFromString(input string) (*exec.Cmd, error) {
	cmd := exec.Command(input)
	return cmd, nil
	// how do I do error handling for exec.Command?
	// e.g. if the input is not a terminal command
	// https://pkg.go.dev/os/exec#Command
}

func ReadInputLoop(input io.Reader) {

}
