package shellspy

import (
	"io"
	"os/exec"
)

func CommandFromString(input string) (*exec.Cmd, error) {
	cmd := exec.Command(input)
	return cmd, nil
}

func ReadInputLoop(input io.Reader) {

}
