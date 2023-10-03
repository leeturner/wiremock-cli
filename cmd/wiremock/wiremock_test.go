package wiremock

import (
	"context"
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
	body, err := wmClient.GetMappings("", 10, 0)
	if err != nil {
		t.Fatal("Error while performing wiremock get mappings", err)
	}
	if body == "" {
		t.Fatal("Expected body to not be empty but got", body)
	}
	if !strings.Contains(body, "4aa0c0b4-a408-4b5e-9325-b3e024a9b674") {
		t.Fatal("Expected body to contain the correct mapping id but got", body)
	}
}

func TestWiremock_GetMappingById(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.GetMappings("0baca68a-0112-4f26-8529-ac12d8eb3720", 10, 0)
	if err != nil {
		t.Fatal("Error while performing wiremock get mappings", err)
	}
	if body == "" {
		t.Fatal("Expected body to not be empty but got", body)
	}
	if !strings.Contains(body, "0baca68a-0112-4f26-8529-ac12d8eb3720") {
		t.Fatal("Expected body to contain the correct mapping id but got", body)
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
	body, err := wmClient.DeleteMappings("")
	if err != nil {
		t.Fatal("Error while performing wiremock get mappings", err)
	}
	if body != "" {
		t.Fatal("Expected body to not be empty but got", body)
	}
}

func TestWiremock_DeleteMappingById(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.DeleteMappings("c15df170-16a4-4d21-8572-ffe6f5f660a3")
	if err != nil {
		t.Fatal("Error while performing wiremock get mappings", err)
	}
	if body != "" {
		t.Fatal("Expected body to not be empty but got", body)
	}
}

func TestWiremock_DeleteMappingByIdNotFound(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.DeleteMappings("ekqg")
	if err != nil {
		t.Fatal("Error while performing wiremock get mappings", err)
	}
	if body != "" {
		t.Fatal("Expected body to be empty but got", body)
	}
}

// Scenario endpoints

func TestWiremock_GetScenarios(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.GetScenarios()
	if err != nil {
		t.Fatal("Error while performing wiremock get scenarios", err)
	}
	if body == "" {
		t.Fatal("Expected body to not be empty but got", body)
	}
	if !strings.Contains(body, "my_scenario") {
		t.Fatal("Expected body to contain the correct scenario name but got", body)
	}
}

// Request endpoints

func TestWiremock_GetRequests(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.GetRequests("", 10)
	if err != nil {
		t.Fatal("Error while performing wiremock get requests", err)
	}
	if body == "" {
		t.Fatal("Expected body to not be empty but got", body)
	}
	if !strings.Contains(body, "45760a03-eebb-4387-ad0d-bb89b5d3d662") {
		t.Fatal("Expected body to contain the correct request but got", body)
	}
}

func TestWiremock_GetRequestById(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.GetRequests("12fb14bb-600e-4bfa-bd8d-be7f12562c9", 10)
	if err != nil {
		t.Fatal("Error while performing wiremock get requests", err)
	}
	if body == "" {
		t.Fatal("Expected body to not be empty but got", body)
	}
	if !strings.Contains(body, "12fb14bb-600e-4bfa-bd8d-be7f12562c9") {
		t.Fatal("Expected body to contain the correct request id but got", body)
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
	body, err := wmClient.DeleteRequests("")
	if err != nil {
		t.Fatal("Error while performing wiremock get mappings", err)
	}
	if body != "" {
		t.Fatal("Expected body to not be empty but got", body)
	}
}

func TestWiremock_DeleteRequestsById(t *testing.T) {
	wmClient, err := initWiremockClient(t)
	if err != nil {
		t.Fatal("Error initialising wiremock container or client", err)
	}
	body, err := wmClient.DeleteRequests("4a343193-a1bf-4b3e-a63b-8c734be18af1")
	if err != nil {
		t.Fatal("Error while performing wiremock get mappings", err)
	}
	if body != "" {
		t.Fatal("Expected body to not be empty but got", body)
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
