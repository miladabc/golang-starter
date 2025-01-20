package config

import (
	"github.com/miladabc/golang-starter/internal/http"
	"github.com/miladabc/golang-starter/internal/log"
	"github.com/miladabc/golang-starter/pkg/mysql"
)

type Config struct {
	Debug  bool              `config:"debug"`
	Log    log.Config        `config:"log" validate:"required"`
	Server http.ServerConfig `config:"server" validate:"required"`
	DB     mysql.Config      `config:"db" validate:"required"`
}
