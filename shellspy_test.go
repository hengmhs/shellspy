package shellspy_test

import (
	"os"
	"os/exec"
	"shellspy"
	"slices"
	"testing"
)

func TestCommandFromStringReturnsExecCmdObject(t *testing.T) {

	t.Parallel()

	cmd := exec.Command("ls")
	want := cmd
	got, err := shellspy.CommandFromString("ls")

	if err != nil {
		t.Error("Err: ", err)
	}

	// Cannot compare *exec.cmd objects directly

	if !slices.Equal(want.Args, got.Args) {
		t.Errorf("Args do not match. want: %#v, got %#v", want, got)
	}
}

func TestCommandFromStringReturnsExecCmdObjectForMultipleInputs(t *testing.T) {

	t.Parallel()

	cmd := exec.Command("echo", "Hello")
	want := cmd
	got, err := shellspy.CommandFromString("echo Hello")

	if err != nil {
		t.Error("Err: ", err)
	}

	// Cannot compare *exec.cmd objects directly

	if !slices.Equal(want.Args, got.Args) {
		t.Errorf("Args do not match. want: %#v, got %#v", want, got)
	}
}

func TestCreateTextFile(t *testing.T) {

	t.Parallel()

	const mockFile = "shellspy_test.txt"

	_, err := os.Stat(mockFile)
	if !os.IsNotExist(err) {
		t.Error("Mock file already exists")
	}

	shellspy.CreateTextFile(mockFile)

	_, err = os.Stat(mockFile)
	if os.IsNotExist(err) {
		t.Error("Mock file was not created")
	}

	os.Remove(mockFile)
}

func TestWriteToTextFile(t *testing.T) {

	const mockFile = "shellspy_test2.txt"

	t.Parallel()

	_, err := os.Stat(mockFile)
	if !os.IsNotExist(err) {
		t.Error("Mock file already exists")
	}

	shellspy.CreateTextFile(mockFile)

	shellspy.WriteToTextFile(mockFile, "echo hello", "hello")

	data, err := os.ReadFile(mockFile)
	if err != nil {
		t.Error("Unable to read mock file")
	}

	got := string(data)
	want := "> echo hello \nhello \n"

	if want != got {
		t.Error("Mock transcript does not match input")
	}

	os.Remove(mockFile)
}
