package todo

import "github.com/labstack/echo/v4"

func (a *App) RegisterRoutes(g *echo.Group) {
	t := g.Group("/todo")

	t.GET("", a.controller.LastTodo)
}
