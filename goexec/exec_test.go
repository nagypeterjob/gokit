package goexec

import (
	"os/exec"
	"reflect"
	"testing"
)

func TestChain(t *testing.T) {

	expected := "test2"

	stdout, stderr, err := Chain(
		exec.Command("echo", "test1", "test2"),
		exec.Command("tr", "\" \"", "\n"),
		exec.Command("grep", "-v", "test1"),
	)

	if err != nil {
		t.Error(stderr)
	}

	if !reflect.DeepEqual(expected, stdout) {
		t.Errorf("TestChain wanted: %v => got: %v", expected, stdout)
	}
}
