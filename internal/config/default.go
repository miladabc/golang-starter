package config

import (
	"time"

	"github.com/miladabc/golang-starter/internal/http"
	"github.com/miladabc/golang-starter/internal/log"
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
}
