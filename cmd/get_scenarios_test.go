package cmd

import (
	"github.com/magiconair/properties/assert"
	"strings"
	"testing"
)

func TestNewGetScenariosCommand(t *testing.T) {
	_, port, err := initContainer(t)
	if err != nil {
		t.Fatal("Error initialising container while running command test", err)
	}
	test := map[string]struct {
		args             []string
		expectedContains string
		expectedError    error
	}{
		"Get all scenarios": {
			args:             []string{"get", "scenarios"},
			expectedContains: "c8d249ec-d86d-48b1-88a8-a660e6848042",
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
