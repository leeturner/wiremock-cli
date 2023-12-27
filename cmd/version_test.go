package cmd

import (
	"strings"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	_, port, err := initContainer(t)
	if err != nil {
		t.Fatal("Error initialising container while running command test", err)
	}
	result, err := ExecuteCommand([]string{"version"}, port)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expectedContains := "3.3.1"
	if !strings.Contains(result, expectedContains) {
		t.Fatal("Unexpected output from command. Expected output to contain: ", expectedContains, " Got: ", result)
	}
}
