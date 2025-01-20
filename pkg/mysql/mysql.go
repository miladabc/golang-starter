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
