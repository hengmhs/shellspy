package shellspy

import (
	"os/exec"
	"slices"
	"testing"
)

func TestCommandFromStringReturnsExecCmdObject(t *testing.T) {
	cmd := exec.Command("ls")
	want := cmd
	got, err := CommandFromString("ls")
	if err != nil {
		t.Error("Err: ", err)
	}
	if !slices.Equal(want.Args, got.Args) {
		t.Errorf("Args do not match. want: %#v, got %#v", want, got)
	}
}
