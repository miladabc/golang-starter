package mysql

import (
	"fmt"
	"net/url"
	"time"
)

type Config struct {
	Host               string        `config:"host" validate:"required"`
	Port               int           `config:"port" validate:"required"`
	User               string        `config:"user" validate:"required"`
	Password           string        `config:"password" validate:"required"`
	DBName             string        `config:"dbname" validate:"required"`
	Charset            string        `config:"charset" validate:"required"`
	Collation          string        `config:"collation" validate:"required"`
	ParseTime          bool          `config:"parse-time"`
	Location           string        `config:"location" validate:"required,timezone"`
	MaxLifeTime        time.Duration `config:"max-life-time" validate:"required"`
	MaxIdleTime        time.Duration `config:"max-idle-time"`
	MaxOpenConnections int           `config:"max-open-connections" validate:"required"`
	MaxIdleConnections int           `config:"max-idle-connections" validate:"required"`
}

func (d Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t&charset=%s&collation=%s&loc=%s&multiStatements=true",
		d.User, d.Password, d.Host, d.Port, d.DBName, d.ParseTime, d.Charset, d.Collation, url.PathEscape(d.Location))
}
