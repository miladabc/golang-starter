package config

import (
	"time"

	"github.com/miladabc/golang-starter/internal/http"
	"github.com/miladabc/golang-starter/internal/log"
	"github.com/miladabc/golang-starter/pkg/mysql"
)

var Default = Config{
	Debug: true,
	Log: log.Config{
		Pretty: true,
		Level:  "trace",
	},
	Server: http.ServerConfig{
		Address:         "0.0.0.0:8080",
		ReadTimeout:     time.Second,
		WriteTimeout:    3 * time.Second,
		IdleTimeout:     2 * time.Minute,
		ShutdownTimeout: 5 * time.Second,
	},
	DB: mysql.Config{
		Host:               "mysql",
		Port:               3306,
		User:               "starter",
		Password:           "starter",
		DBName:             "starter",
		Charset:            "utf8mb4",
		Collation:          "utf8mb4_unicode_ci",
		ParseTime:          true,
		Location:           "Asia/Tehran",
		MaxLifeTime:        5 * time.Minute,
		MaxIdleTime:        0 * time.Second,
		MaxOpenConnections: 10,
		MaxIdleConnections: 5,
	},
}
