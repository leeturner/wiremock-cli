package cmd

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDeleteRequestsCommand(t *testing.T) {
	_, port, err := initContainer(t)
	if err != nil {
		t.Fatal("Error initialising container while running command test", err)
	}
	test := map[string]struct {
		args           []string
		expectedOutput string
		expectedError  error
	}{
		"Delete all requests": {
			args:           []string{"requests", "delete"},
			expectedOutput: "",
			expectedError:  nil,
		},
		"Delete requests by ID": {
			args:           []string{"requests", "delete", "--id", "4a343193-a1bf-4b3e-a63b-8c734be18af1"},
			expectedOutput: "",
			expectedError:  nil,
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			result, err := ExecuteCommand(tc.args, port)
			assert.Equal(t, result, tc.expectedOutput)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
