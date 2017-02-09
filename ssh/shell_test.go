package ssh

import (
	"testing"
)

// TestRunCommand - test run command
func TestRunCommand(t *testing.T) {
	cmds := []string{
		"pwd",
		"ls",
	}

	for _, cmd := range cmds {
		out, err := RunCommand(cmd)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(out)
	}
}

// TestRunShellCommand - test run shell command
func TestRunShellCommand(t *testing.T) {
	cmds := []string{
		"pwd",
		"free -g",
		"head -n 100 /etc/init.d/ssh | grep \"END\"",
	}

	for _, cmd := range cmds {
		out, err := RunShellCommand(cmd)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(out)
	}
}
