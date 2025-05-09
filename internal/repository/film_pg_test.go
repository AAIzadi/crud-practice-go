package repository_test

import (
	"context"
	appconfig "crud-practice-go/internal/config"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
	"time"
)

func setupPostgresContainer(ctx context.Context, t *testing.T) (*pgxpool.Pool, func(), error) {

	env := "test"
	cfg, _ := appconfig.GetConfig(env)

	t.Helper()

	req := testcontainers.ContainerRequest{
		Image:        "postgres",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       cfg.Postgres.Database,
			"POSTGRES_USER":     cfg.Postgres.Username,
			"POSTGRES_PASSWORD": cfg.Postgres.Password,
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections").
			WithOccurrence(2).
			WithStartupTimeout(30 * time.Second),
	}

	pgContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start container: %v", err)
	}

	host, err := pgContainer.Host(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get container host: %v", err)
	}

	port, err := pgContainer.MappedPort(ctx, "5432")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get container port: %v", err)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.Username, cfg.Postgres.Password, host, port.Port(), cfg.Postgres.Database)

	var pool *pgxpool.Pool
	maxAttempts := 5
	for i := 0; i < maxAttempts; i++ {
		pool, err = pgxpool.New(ctx, connStr)
		if err == nil {
			break
		}
		time.Sleep(time.Second * time.Duration(i+1))
	}
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	cleanup := func() {
		pool.Close()
		pgContainer.Terminate(ctx)
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
