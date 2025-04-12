package shellspy

import (
	"os"
	"os/exec"
	"slices"
	"testing"
)

// why don't I have to import shellspy.CommandFromString?
// is it because i'm using the package shellspy?

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

// func TestReadInputLoop(t *testing.T) {
// 	// os.Stdin is of type *os.File that implements io.Reader interface
// 	// what does that mean?? basically we can pass it to bufio.NewScanner

// 	// simulate user input
// 	input := "Hello\nWorld\nexit\n"
// 	r := strings.NewReader(input)
// 	ReadInputLoop(r)
// }

func TestCreateShellSpy(t *testing.T) {

	// Check that mock file does not exist
	_, err := os.Stat("shellspy_test.txt")
	if !os.IsNotExist(err) {
		t.Error("Mock file already exists")
	}

	CreateTextFile("shellspy_test.txt")

	_, err = os.Stat("shellspy_test.txt")
	if os.IsNotExist(err) {
		t.Error("Mock file was not created")
	}

	// Clean up
	os.Remove("shellspy_test.txt")
}
