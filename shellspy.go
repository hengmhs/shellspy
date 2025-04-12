package shellspy

import (
	"io"
	"os/exec"
	"strings"
)

func CommandFromString(input string) (*exec.Cmd, error) {
	inputs := strings.Split(input, " ")
	cmd := exec.Command(inputs[0], inputs[1:]...)
	return cmd, nil
	// how do I do error handling for exec.Command?
	// e.g. if the input is not a terminal command
	// https://pkg.go.dev/os/exec#Command
}

func ReadInputLoop(input io.Reader) {

}
