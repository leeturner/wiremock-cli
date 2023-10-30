package cmd

import (
	"testing"
)

func TestResetCommand(t *testing.T) {
	_, port, err := initContainer(t)
	if err != nil {
		t.Fatal("Error initialising container while running command test", err)
	}
	result, err := ExecuteCommand([]string{"reset"}, port)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := "Wiremock server reset\n"
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}
