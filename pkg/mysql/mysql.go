package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Connect(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("openning connection to mysql: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetConnMaxLifetime(cfg.MaxLifeTime)
	db.SetConnMaxIdleTime(cfg.MaxIdleTime)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("pinging mysql: %w", err)
	}

	return db, nil
}

func CreateDB(cfg Config) error {
	cfgWithoutDBName := cfg
	cfgWithoutDBName.DBName = ""

	db, err := Connect(cfgWithoutDBName)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET %s COLLATE %s",
		cfg.DBName, cfg.Charset, cfg.Collation,
	)

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("creating database: %w", err)
	}

	return nil
}

func DropDB(cfg Config) error {
	cfgWithoutDBName := cfg
	cfgWithoutDBName.DBName = ""

	db, err := Connect(cfgWithoutDBName)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", cfg.DBName)

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("dropping database: %w", err)
	}

	return nil
}
