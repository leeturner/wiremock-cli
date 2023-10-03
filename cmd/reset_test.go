package cmd

import (
	"testing"
)

func TestResetCommand(t *testing.T) {
	result, err := ExecuteCommand([]string{"reset"}, t)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := "Wiremock server reset\n"
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}
