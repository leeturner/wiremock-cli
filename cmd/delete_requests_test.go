package cmd

import (
	"testing"
)

func TestDeleteRequestsCommand(t *testing.T) {
	result, err := ExecuteCommand([]string{"delete", "requests"}, t)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := ""
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}

func TestDeleteRequestsCommandWithId(t *testing.T) {
	result, err := ExecuteCommand([]string{"delete", "requests", "--id", "4a343193-a1bf-4b3e-a63b-8c734be18af1"}, t)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := ""
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}
