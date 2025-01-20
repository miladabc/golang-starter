package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Server struct {
	cfg    ServerConfig
	server *http.Server
}

type ServerConfig struct {
	Address         string        `config:"address" validate:"required,hostname_port"`
	ReadTimeout     time.Duration `config:"read-timeout" validate:"required"`
	WriteTimeout    time.Duration `config:"write-timeout" validate:"required"`
	IdleTimeout     time.Duration `config:"idle-timeout" validate:"required"`
	ShutdownTimeout time.Duration `config:"shutdown-timeout" validate:"required"`
}

func NewServer(cfg ServerConfig, handler http.Handler) *Server {
	return &Server{
		cfg: cfg,
		server: &http.Server{
			Addr:              cfg.Address,
			ReadTimeout:       cfg.ReadTimeout,
			ReadHeaderTimeout: cfg.ReadTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       cfg.IdleTimeout,
			Handler:           handler,
		},
	}
}

func (s *Server) Start() {
	log.Info().Msgf("http server listening on `%s`", s.server.Addr)

	err := s.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error().Err(err).Msg("starting http server")
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	log.Info().Msg("shutting down http server...")

	deadline, cancel := context.WithTimeout(ctx, s.cfg.ShutdownTimeout)
	defer cancel()

	if err := s.server.Shutdown(deadline); err != nil {
		log.Error().Msgf("shutting down http server: %s", err)
	}
}
