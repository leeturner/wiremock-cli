package cmd

import (
	"github.com/magiconair/properties/assert"
	"strings"
	"testing"
)

func TestGetRequestsCommand(t *testing.T) {
	_, port, err := initContainer(t)
	if err != nil {
		t.Fatal("Error initialising container while running command test", err)
	}
	test := map[string]struct {
		args             []string
		expectedContains string
		expectedError    error
	}{
		"Get all requests": {
			args:             []string{"get", "requests"},
			expectedContains: "45760a03-eebb-4387-ad0d-bb89b5d3d662",
			expectedError:    nil,
		},
		"Get request by ID": {
			args:             []string{"get", "requests", "--id", "12fb14bb-600e-4bfa-bd8d-be7f12562c9"},
			expectedContains: "12fb14bb-600e-4bfa-bd8d-be7f12562c9",
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

func TestGetRequestsCommandWithIdNotFound(t *testing.T) {
	_, port, err := initContainer(t)
	if err != nil {
		t.Fatal("Error initialising container while running command test", err)
	}
	result, err := ExecuteCommand([]string{"get", "requests", "--id", "mxp2"}, port)
	if err != nil {
		t.Fatal("Error running command test", err)
	}
	expected := ""
	if result != expected {
		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	}
}
