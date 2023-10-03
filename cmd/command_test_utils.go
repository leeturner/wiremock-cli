package cmd

import (
	"bytes"
	"context"
	. "github.com/wiremock/wiremock-testcontainers-go"
	"testing"
)

const testAdminPrefix = "__test_admin"

func initContainer(t *testing.T) (host string, port string, err error) {
	ctx := context.Background()
	// Create Container - use the nightly build of wiremock
	container, err := RunContainer(ctx,
		WithImage("wiremock/wiremock:nightly"),
		WithMappingFile("wiremock-admin", "../stubs/wiremock-admin-stubs.json"),
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

func ExecuteCommand(args []string, t *testing.T) (string, error) {
	_, port, err := initContainer(t)
	if err != nil {
		return "", err
	}
	all := append(args, []string{"--port", port, "--admin-prefix", testAdminPrefix}...)
	actual := new(bytes.Buffer)
	rootCmd := NewRootCmd()
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs(all)
	err = rootCmd.Execute()
	if err != nil {
		return "", err
	}
	return actual.String(), nil
}
