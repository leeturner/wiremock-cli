package cmd

import (
	"github.com/magiconair/properties/assert"
	"strings"
	"testing"
)

func TestGetRecordingsStatus(t *testing.T) {
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
			args:             []string{"recordings", "status"},
			expectedContains: "Stopped",
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
