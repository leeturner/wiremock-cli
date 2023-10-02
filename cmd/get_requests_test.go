package cmd

import (
	"strings"
	"testing"
)

func TestGetRequestsCommand(t *testing.T) {
	result, err := ExecuteCommand([]string{"get", "requests"})
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := "45760a03-eebb-4387-ad0d-bb89b5d3d662"
	if !strings.Contains(result, expected) {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Within: ", result)
	}
	if !strings.Contains(result, "\"total\" : 9") {
		t.Fatal("Unexpected output from command. Expected: ", "\"total\" : 2", " Within: ", result)
	}
}

func TestGetRequestsCommandWithId(t *testing.T) {
	result, err := ExecuteCommand([]string{"get", "requests", "--id", "12fb14bb-600e-4bfa-bd8d-be7f12562c9"})
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := "12fb14bb-600e-4bfa-bd8d-be7f12562c9"
	if !strings.Contains(result, expected) {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}

func TestGetRequestsCommandWithIdNotFound(t *testing.T) {
	result, err := ExecuteCommand([]string{"get", "requests", "--id", "mxp2"})
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := ""
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}
