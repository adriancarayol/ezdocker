package cli

import (
	"github.com/adriancarayol/ezdocker/internal/docker"
	"testing"
)

func TestInit(t *testing.T) {
	dockerClient := docker.New()
	printContainersCmd := docker.PrintContainersCommand{Docker: dockerClient}
	expected := []Command{
		{CommandName: list, Handler: printContainersCmd.Handle},
	}

	Init()

	if len(expected) != len(options) {
		t.Fatalf("Failed. len(expected) = %d, len(options) = %d", len(expected), len(options))
	}

	matchingCommands := 0

	for _, e := range expected {
		for _, o := range options {
			if e.CommandName == o.CommandName {
				matchingCommands += 1
			}
		}
	}

	if matchingCommands != len(expected) {
		t.Fatalf("Failed. Matching commands: %d, expected: %d", matchingCommands, len(expected))
	}
}
