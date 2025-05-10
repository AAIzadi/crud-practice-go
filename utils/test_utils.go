package utils

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func FindProjectRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd, nil
		}

		parent := filepath.Dir(wd)
		if parent == wd {
			return "", os.ErrNotExist
		}
		wd = parent
	}
}

func LoadSQLFile(t *testing.T, db *pgxpool.Pool, base string, file string) {
	content, err := os.ReadFile(filepath.Join(base, file))
	require.NoError(t, err)

	_, err = db.Exec(context.Background(), string(content))
	require.NoError(t, err)
}
