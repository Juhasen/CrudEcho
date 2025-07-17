package testhelpers

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreatePostgresContainer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pgContainer, err := CreatePostgresContainer(ctx)
	require.NoError(t, err, "failed to start postgres container")
	require.NotNil(t, pgContainer)
	require.NotEmpty(t, pgContainer.ConnectionString)

	t.Logf("Postgres container started with connection string: %s", pgContainer.ConnectionString)

	// Terminate container after test to clean up
	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Logf("failed to terminate container: %v", err)
		}
	}()
}
