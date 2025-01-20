package mysql

import (
	"database/sql"
	"io/fs"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func NewMigration(s fs.FS, path string, db *sql.DB, dbname string) (*migrate.Migrate, error) {
	source, err := iofs.New(s, path)
	if err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{
		DatabaseName:    dbname,
		MigrationsTable: "schema_migrations",
	})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithInstance("iofs", source, dbname, driver)
}
