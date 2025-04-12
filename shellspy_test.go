package shellspy

import (
	"os/exec"
	"slices"
	"strings"
	"testing"
)

func TestCommandFromStringReturnsExecCmdObject(t *testing.T) {
	cmd := exec.Command("ls")
	want := cmd
	got, err := CommandFromString("ls")

	if err != nil {
		t.Error("Err: ", err)
	}

	// Cannot compare *exec.cmd objects directly

	if !slices.Equal(want.Args, got.Args) {
		t.Errorf("Args do not match. want: %#v, got %#v", want, got)
	}
}

func TestCommandFromStringReturnsExecCmdObjectForMultipleInputs(t *testing.T) {
	cmd := exec.Command("echo", "Hello")
	want := cmd
	got, err := CommandFromString("echo Hello")

	if err != nil {
		t.Error("Err: ", err)
	}

	// Cannot compare *exec.cmd objects directly

	if !slices.Equal(want.Args, got.Args) {
		t.Errorf("Args do not match. want: %#v, got %#v", want, got)
	}
}

func TestReadInputLoop(t *testing.T) {
	// os.Stdin is of type *os.File that implements io.Reader interface
	// what does that mean?? basically we can pass it to bufio.NewScanner

	// simulate user input
	input := "Hello\nWorld\nexit\n"
	r := strings.NewReader(input)
	ReadInputLoop(r)
}
