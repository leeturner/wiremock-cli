package cmd

import (
	"github.com/magiconair/properties/assert"
	"strings"
	"testing"
)

func TestGetMappingsCommand(t *testing.T) {
	_, port, err := initContainer(t)
	if err != nil {
		t.Fatal("Error initialising container while running command test", err)
	}
	test := map[string]struct {
		args             []string
		expectedContains string
		expectedError    error
	}{
		"Get all mappings": {
			args:             []string{"mappings", "get"},
			expectedContains: "4aa0c0b4-a408-4b5e-9325-b3e024a9b674",
			expectedError:    nil,
		},
		"Get mapping by ID": {
			args:             []string{"mappings", "get", "--id", "0baca68a-0112-4f26-8529-ac12d8eb3720"},
			expectedContains: "0baca68a-0112-4f26-8529-ac12d8eb3720",
			expectedError:    nil,
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			result, err := ExecuteCommand(tc.args, port)
			assert.Equal(t, true, strings.Contains(result, tc.expectedContains))
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetMappingCommandWithIdNotFound(t *testing.T) {
	_, port, err := initContainer(t)
	if err != nil {
		t.Fatal("Error initialising container while running command test", err)
	}
	result, err := ExecuteCommand([]string{"mappings", "get", "--id", "e148"}, port)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := ""
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}
