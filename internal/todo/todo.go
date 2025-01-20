package todo

import (
	echo "github.com/labstack/echo/v4"
	"github.com/miladabc/golang-starter/internal/http"
)

func Init(router *http.Router) {
	g := router.V1.Group("/todo")

	g.GET("", func(c echo.Context) error {
		return c.String(200, "todo")
	})
}
