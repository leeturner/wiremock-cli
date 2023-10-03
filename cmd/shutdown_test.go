package cmd

import (
	"testing"
)

func TestShutdownCommand(t *testing.T) {
	result, err := ExecuteCommand([]string{"shutdown"}, t)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := "Wiremock server shutdown\n"
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}
