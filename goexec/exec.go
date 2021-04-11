package goexec

//
// The following file is based on this repository:
// https://gist.github.com/kylelemons/1525278
//

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
)

// Output runs single command
//
// Returns stdout as string & without terminating "\n" character
func Output(ctx context.Context, command string, args ...string) (string, error) {
	out, err := exec.CommandContext(ctx, command, args...).Output()
	return strings.TrimSuffix(string(out), "\n"), err
}

// Run runs single command
//
// returns stdout / stderr as byte arrays and error
func Run(ctx context.Context, command string, args ...string) ([]byte, []byte, error) {
	var (
		stdout = &bytes.Buffer{}
		stderr = &bytes.Buffer{}
	)
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return stdout.Bytes(), stderr.Bytes(), cmd.Run()
}

// RunWith takes stdout / stderr as bytes.Buffer pointers
func RunWith(ctx context.Context, stdout *bytes.Buffer, stderr *bytes.Buffer, command string, args ...string) error {
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return cmd.Run()
}

// Enables to run piped commands
//
// Returns stdout as string & without terminating "\n" character
//
// Example usage:
//
//	stdout, stderr, err := Chain(
//		exec.Command("echo", "test1", "test2"),
//		exec.Command("tr", "\" \"", "\n"),
//		exec.Command("grep", "-v", "test1"),
//	)
//
// The above snippet is equivalent to:
//
// $ echo "test1" "test2" | tr " " "\n" |grep -v "test1"
//
// $ test2
func Chain(commands ...*exec.Cmd) (string, string, error) {

	var (
		length = len(commands)

		// Create buffers for each commands to avoid race condition
		stdout = make([]bytes.Buffer, length)
		stderr = make([]bytes.Buffer, length)
	)

	if length < 1 {
		return "", "", nil
	}

	last := length - 1
	for i, cmd := range commands[:last] {
		var err error
		if commands[i+1].Stdin, err = cmd.StdoutPipe(); err != nil {
			return "", "", err
		}

		cmd.Stderr = &stderr[i]
	}

	commands[last].Stdout, commands[last].Stderr = &stdout[last], &stderr[last]

	for i, cmd := range commands {
		if err := cmd.Start(); err != nil {
			return stdout[i].String(), stderr[i].String(), err
		}
	}

	for i, cmd := range commands {
		if err := cmd.Wait(); err != nil {
			return stdout[i].String(), stderr[i].String(), err
		}
	}

	return strings.TrimSpace(stdout[last].String()), stderr[last].String(), nil
}
