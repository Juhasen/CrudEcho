package testhelpers

/*
Commands to run to testcontainers work:
- export TESTCONTAINERS_RYUK_DISABLED=true
- export DOCKER_HOST="unix://${HOME}/.colima/default/docker.sock"
- export TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE=/var/run/docker.sock
- export TESTCONTAINERS_HOST_OVERRIDE=$(colima ls -j | jq -r '.address')
- source ~/.zshrc
- docker context use colima
- go test -v
*/

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"
)

type PostgresContainer struct {
	testcontainers.Container
	ConnectionString string
}

func CreatePostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:16-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "testcrud",
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "123",
		},
		WaitingFor: wait.ForAll(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(time.Minute),
		),
	}

	ctr, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	host := "localhost"

	port, err := ctr.MappedPort(ctx, "5432")
	if err != nil {
		err := ctr.Terminate(ctx)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	connStr := fmt.Sprintf(
		"postgres://postgres:123@%s:%s/testcrud?sslmode=disable",
		host, port.Port(),
	)

	return &PostgresContainer{
		Container:        ctr,
		ConnectionString: connStr,
	}, nil
}
