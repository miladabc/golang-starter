package docs

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	httpSwagger "github.com/swaggo/http-swagger"
)

//go:embed openapi.yaml
var fs embed.FS

func RegisterRoutes(g *echo.Group) {
	g.GET("/swagger/*", echo.WrapHandler(httpSwagger.Handler(
		httpSwagger.URL("/docs/openapi.yaml"),
	)))

	g.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	g.StaticFS("/docs", fs)
}
