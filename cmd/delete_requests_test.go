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
			args:           []string{"delete", "requests"},
			expectedOutput: "",
			expectedError:  nil,
		},
		"Delete requests by ID": {
			args:           []string{"delete", "requests", "--id", "4a343193-a1bf-4b3e-a63b-8c734be18af1"},
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

	//result, err := ExecuteCommand([]string{"delete", "requests"}, t)
	//if err != nil {
	//	t.Fatal("Error running command test", err)
	//}
	//expected := ""
	//if result != expected {
	//	t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
	//}
}

//func TestDeleteRequestsCommandWithId(t *testing.T) {
//	result, err := ExecuteCommand([]string{"delete", "requests", "--id", "4a343193-a1bf-4b3e-a63b-8c734be18af1"}, t)
//	if err != nil {
//		t.Fatal("Error running command test", err)
//	}
//	expected := ""
//	if result != expected {
//		t.Fatal("Unexpected output from command. Expected: ", expected, " Got: ", result)
//	}
//}
