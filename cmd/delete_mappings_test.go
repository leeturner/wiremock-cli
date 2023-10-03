package cmd

import (
	"testing"
)

func TestDeleteMappingsCommand(t *testing.T) {
	result, err := ExecuteCommand([]string{"delete", "mappings"}, t)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := ""
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}

func TestDeleteMappingsCommandWithId(t *testing.T) {
	result, err := ExecuteCommand([]string{"delete", "mappings", "--id", "c15df170-16a4-4d21-8572-ffe6f5f660a3"}, t)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := ""
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}

func TestDeleteMappingCommandWithIdNotFound(t *testing.T) {
	result, err := ExecuteCommand([]string{"delete", "mappings", "--id", "ekqg"}, t)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := ""
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}
