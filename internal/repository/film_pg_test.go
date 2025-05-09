package repository_test

import (
	"context"
	appconfig "crud-practice-go/internal/config"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"testing"
	"time"
)

func setupPostgresContainer(ctx context.Context, t *testing.T) (*pgxpool.Pool, func(), error) {

	cfg, _ := appconfig.GetConfig("test")
	postgresConfig := cfg.Postgres

	t.Helper()

	ctr, err := postgres.Run(ctx, "postgres",
		postgres.WithDatabase(postgresConfig.Database),
		postgres.WithUsername(postgresConfig.Username),
		postgres.WithPassword(postgresConfig.Password),
		postgres.BasicWaitStrategies(),
	)

	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	connStr, err := ctr.ConnectionString(ctx, "sslmode=disable", "application_name=test")
	require.NoError(t, err)

	id, err := ctr.MappedPort(ctx, "5432/tcp")
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&application_name=test",
		postgresConfig.Username,
		postgresConfig.Password,
		"localhost", id.Port(),
		postgresConfig.Database),
		connStr)

	var pool *pgxpool.Pool
	maxAttempts := 5
	for i := 0; i < maxAttempts; i++ {
		pool, err = pgxpool.New(ctx, connStr)
		if err == nil {
			break
		}
		time.Sleep(time.Second * time.Duration(i+1))
	}

	require.NoError(t, err)
	require.NotNil(t, pool)

	cleanup := func() {
		pool.Close()
		ctr.Terminate(ctx)
	}

	return pool, cleanup, nil
}

func createFilmTable(ctx context.Context, db *pgxpool.Pool, t *testing.T) {
	_, err := db.Exec(ctx, `
        CREATE TABLE film (
            film_id SERIAL PRIMARY KEY,
            title TEXT,
            description TEXT,
            release_year INT,
            language_id INT,
            rating TEXT
        );
    `)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFilmRepository_GetAll(t *testing.T) {
	ctx := context.Background()
	pool, cleanup, err := setupPostgresContainer(ctx, t)
	assert.NoError(t, err)
	defer cleanup()

	createFilmTable(ctx, pool, t)
}
