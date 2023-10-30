package wiremock

import (
	"context"
	"github.com/magiconair/properties/assert"
	. "github.com/wiremock/wiremock-testcontainers-go"
	"strings"
	"testing"
)

const testAdminPrefix = "__test_admin"

func initContainer(t *testing.T) (host string, port string, err error) {
	ctx := context.Background()
	// Create Container - use the nightly build of wiremock
	container, err := RunContainer(ctx,
		WithImage("wiremock/wiremock:nightly"),
		WithMappingFile("wiremock-admin", "../../stubs/wiremock-admin-stubs.json"),
	)
	if err != nil {
		return "", "", err
	}
	containerHost, err := container.Host(ctx)
	if err != nil {
		return "", "", err
	}
	ports, err := container.Ports(ctx)
	if err != nil {
		return "", "", err
	}
	containerPort := ports["8080/tcp"][0].HostPort

	defer t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})
	return containerHost, containerPort, nil
}

func initWiremockClient(t *testing.T) (wmClient *Wiremock, err error) {
	host, port, err := initContainer(t)
	if err != nil {
		return nil, err
	}
	wmClient = Init("http://"+host, port)
	wmClient.WithAdminPrefix(testAdminPrefix)
	return wmClient, nil
}

func TestInitClientIsNotNil(t *testing.T) {
	wmClient := Init("http://localhost", "8080")
	if wmClient == nil {
		t.Error("Expected wiremock client to not be nil")
	}
}

// Mapping endpoints

func TestWiremock_GetMappings(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	test := map[string]struct {
		id               string
		limit            int
		offset           int
		expectedContains string
		expectedError    error
	}{
		"Get all mappings": {
			id:               "",
			limit:            10,
			offset:           0,
			expectedContains: "4aa0c0b4-a408-4b5e-9325-b3e024a9b674",
			expectedError:    nil,
		},
		"Get mappings by id": {
			id:               "0baca68a-0112-4f26-8529-ac12d8eb3720",
			limit:            10,
			offset:           0,
			expectedContains: "0baca68a-0112-4f26-8529-ac12d8eb3720",
			expectedError:    nil,
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			actual, actualErr := wmClient.GetMappings(tc.id, tc.limit, tc.offset)
			assert.Equal(t, true, strings.Contains(actual, tc.expectedContains))
			assert.Equal(t, tc.expectedError, actualErr)
		})
	}
}

func TestWiremock_GetMappingByIdNotFound(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.GetMappings("e148", 10, 0)
	if err != nil {
		t.Fatal("Error while performing wiremock get mappings", err)
	}
	if body != "" {
		t.Fatal("Expected body to be empty but got", body)
	}
}

func TestWiremock_DeleteMappings(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	test := map[string]struct {
		id            string
		expected      string
		expectedError error
	}{
		"Delete all mappings": {
			id:            "",
			expected:      "",
			expectedError: nil,
		},
		"Delete mappings by id": {
			id:            "c15df170-16a4-4d21-8572-ffe6f5f660a3",
			expected:      "",
			expectedError: nil,
		},
		"Delete mappings by id not found": {
			id:            "ekqg",
			expected:      "",
			expectedError: nil,
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			actual, actualErr := wmClient.DeleteMappings(tc.id)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedError, actualErr)
		})
	}
}

// Scenario endpoints

func TestWiremock_GetScenarios(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	test := map[string]struct {
		id               string
		limit            int
		offset           int
		expectedContains string
		expectedError    error
	}{
		"Get all mappings": {
			expectedContains: "my_scenario",
			expectedError:    nil,
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			actual, actualErr := wmClient.GetScenarios()
			assert.Equal(t, true, strings.Contains(actual, tc.expectedContains))
			assert.Equal(t, tc.expectedError, actualErr)
		})
	}
}

// Request endpoints

func TestWiremock_GetRequests(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	test := map[string]struct {
		id               string
		limit            int
		expectedContains string
		expectedError    error
	}{
		"Get all requests": {
			id:               "",
			limit:            10,
			expectedContains: "45760a03-eebb-4387-ad0d-bb89b5d3d662",
			expectedError:    nil,
		},
		"Get requests by id": {
			id:               "12fb14bb-600e-4bfa-bd8d-be7f12562c9",
			limit:            10,
			expectedContains: "12fb14bb-600e-4bfa-bd8d-be7f12562c9",
			expectedError:    nil,
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			actual, actualErr := wmClient.GetRequests(tc.id, tc.limit)
			assert.Equal(t, true, strings.Contains(actual, tc.expectedContains))
			assert.Equal(t, tc.expectedError, actualErr)
		})
	}
}

func TestWiremock_GetRequestByIdNotFound(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.GetRequests("mxp2", 10)
	if err != nil {
		t.Fatal("Error while performing wiremock get requests", err)
	}
	if body != "" {
		t.Fatal("Expected body to be empty but got", body)
	}
}

func TestWiremock_DeleteRequests(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	test := map[string]struct {
		id            string
		expected      string
		expectedError error
	}{
		"Delete all requests": {
			id:            "",
			expected:      "",
			expectedError: nil,
		},
		"Delete request by id": {
			id:            "4a343193-a1bf-4b3e-a63b-8c734be18af1",
			expected:      "",
			expectedError: nil,
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			actual, actualErr := wmClient.DeleteRequests(tc.id)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedError, actualErr)
		})
	}
}

// System endpoints

func TestWiremock_Shutdown(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.Shutdown()
	if err != nil {
		t.Fatal("Error while performing wiremock shutdown", err)
	}
	if body != "" {
		t.Fatal("Expected body to be empty but got", body)
	}
}

func TestWiremock_Reset(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.Reset()
	if err != nil {
		t.Fatal("Error while performing wiremock reset", err)
	}
	if body != "" {
		t.Fatal("Expected body to be empty but got", body)
	}
}
