package app

import (
	"context"
	"errors"
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/miladabc/golang-starter/internal/container"
	"github.com/miladabc/golang-starter/internal/migrations"
	"github.com/miladabc/golang-starter/pkg/mysql"
	"github.com/rs/zerolog/log"
)

func RunDBMigrations() error {
	ctx := context.Background()
	c := container.New()
	defer c.Shutdown(ctx)

	err := c.Init()
	if err != nil {
		return err
	}

	m, err := mysql.NewMigration(mysql.MigrationConfig{
		FS:     migrations.FS,
		Path:   migrations.Path,
		DB:     c.DB.DB,
		DBName: c.Config.DB.DBName,
	})
	if err != nil {
		return fmt.Errorf("preparing migration: %w", err)
	}

	log.Info().Msg("Migrations in progress...")

	err = m.Up()

	switch {
	case err == nil:
		log.Info().Msg("migrations applied successfully")
	case errors.Is(err, migrate.ErrNoChange):
		log.Info().Msg("no migrations to apply")
	default:
		return fmt.Errorf("running migrations: %w", err)
	}

	return nil
}
