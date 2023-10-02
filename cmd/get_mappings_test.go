package cmd

import (
	"strings"
	"testing"
)

func TestGetMappingsCommand(t *testing.T) {
	result, err := ExecuteCommand([]string{"get", "mappings"})
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := "4aa0c0b4-a408-4b5e-9325-b3e024a9b674"
	if !strings.Contains(result, expected) {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Within: ", result)
	}
	if !strings.Contains(result, "\"total\" : 2") {
		t.Fatal("Unexpected output from command. Expected: ", "\"total\" : 2", " Within: ", result)
	}
}

func TestGetMappingsCommandWithId(t *testing.T) {
	result, err := ExecuteCommand([]string{"get", "mappings", "0baca68a-0112-4f26-8529-ac12d8eb3720"})
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := "0baca68a-0112-4f26-8529-ac12d8eb3720"
	if !strings.Contains(result, expected) {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}

func TestGetMappingCommandWithIdNotFound(t *testing.T) {
	result, err := ExecuteCommand([]string{"get", "mappings", "e148"})
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := ""
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}
