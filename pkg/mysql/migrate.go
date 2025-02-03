package mysql

import (
	"database/sql"
	"io/fs"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type MigrationConfig struct {
	FS     fs.FS
	Path   string
	DB     *sql.DB
	DBName string
}

func NewMigration(cfg MigrationConfig) (*migrate.Migrate, error) {
	source, err := iofs.New(cfg.FS, cfg.Path)
	if err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(cfg.DB, &mysql.Config{
		DatabaseName:    cfg.DBName,
		MigrationsTable: "schema_migrations",
	})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithInstance("iofs", source, cfg.DBName, driver)
}
