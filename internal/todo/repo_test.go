package todo

import (
	"context"
	"testing"
	"time"

	"github.com/miladabc/golang-starter/internal/config/configtest"
	"github.com/miladabc/golang-starter/internal/migrations"
	"github.com/miladabc/golang-starter/pkg/mysql/mysqltest"
	"github.com/stretchr/testify/require"
)

func TestMySQLRepoCreate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg := configtest.New(t)
	db := mysqltest.ConnectAndMigrate(t, cfg.DB, migrations.FS, ".")
	repo := NewRepository(db)
	expectedTodo := Todo{
		ID:          1,
		Description: "desc",
		DueDate:     time.Now(),
	}

	actualTodo, err := repo.Create(ctx, expectedTodo.Description, expectedTodo.DueDate)
	require.NoError(t, err)
	require.Equal(t, expectedTodo, actualTodo)

	lastTodo, err := repo.FindLast(ctx)
	require.NoError(t, err)
	require.Equal(t, expectedTodo.ID, lastTodo.ID)
	require.Equal(t, expectedTodo.Description, lastTodo.Description)
	require.True(t, expectedTodo.DueDate.Equal(actualTodo.DueDate))
}

func TestMySQLRepoFindLast_NotFound(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg := configtest.New(t)
	db := mysqltest.ConnectAndMigrate(t, cfg.DB, migrations.FS, ".")
	repo := NewRepository(db)

	_, err := repo.FindLast(ctx)
	require.ErrorIs(t, err, ErrTodoNotFound)
}

func TestMySQLRepoFindLast_Err(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg := configtest.New(t)
	db := mysqltest.ConnectAndMigrate(t, cfg.DB, migrations.FS, ".")
	repo := NewRepository(db)

	db.Close()

	_, err := repo.FindLast(ctx)
	require.ErrorContains(t, err, "database is closed")
}
