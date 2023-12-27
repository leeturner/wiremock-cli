package cmd

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDeleteMappingsCommands(t *testing.T) {
	_, port, err := initContainer(t)
	if err != nil {
		t.Fatal("Error initialising container while running command test", err)
	}
	test := map[string]struct {
		args           []string
		expectedOutput string
		expectedError  error
	}{
		"Delete all mappings": {
			args:           []string{"mappings", "delete"},
			expectedOutput: "",
			expectedError:  nil,
		},
		"Delete mapping by ID": {
			args:           []string{"mappings", "delete", "--id", "c15df170-16a4-4d21-8572-ffe6f5f660a3"},
			expectedOutput: "",
			expectedError:  nil,
		},
		"Delete mapping by ID not found": {
			args:           []string{"mappings", "delete", "--id", "ekqg"},
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
