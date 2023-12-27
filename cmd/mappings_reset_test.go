package cmd

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewResetMappingsCommand(t *testing.T) {
	_, port, err := initContainer(t)
	if err != nil {
		t.Fatal("Error initialising container while running command test", err)
	}
	test := map[string]struct {
		args          []string
		expected      string
		expectedError error
	}{
		"Reset mappings": {
			args:          []string{"mappings", "reset"},
			expected:      "",
			expectedError: nil,
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			result, err := ExecuteCommand(tc.args, port)
			assert.Equal(t, result, tc.expected)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}
