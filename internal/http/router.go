package http

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/miladabc/golang-starter/pkg/middleware"
	"github.com/miladabc/golang-starter/pkg/validator"
)

type Router struct {
	r    *echo.Echo
	Root *echo.Group
	V1   *echo.Group
}

type RouterConfig struct {
	Debug bool `config:"debug"`
}

func NewRouter(cfg RouterConfig) *Router {
	router := newEchoRouter(cfg)

	return &Router{
		r:    router,
		Root: router.Group(""),
		V1:   router.Group("/api/v1"),
	}
}

func (r *Router) Handler() http.Handler {
	return r.r
}

func newEchoRouter(cfg RouterConfig) *echo.Echo {
	router := echo.New()
	router.Debug = cfg.Debug
	router.Validator = validator.NewEchoValidator()

	router.Pre(echoMiddleware.RemoveTrailingSlash())
	router.Use(echoMiddleware.Recover())
	router.Use(middleware.AccessRequestLogger())
	router.Use(echoMiddleware.CORS())

	return router
}
