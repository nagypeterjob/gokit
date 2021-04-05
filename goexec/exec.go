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

// Runs single command
//
// Returns stdout as string & without terminating "\n" character
func Run(ctx context.Context, command string, args ...string) (string, error) {
	out, err := exec.CommandContext(ctx, command, args...).Output()
	return strings.TrimSuffix(string(out), "\n"), err
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
