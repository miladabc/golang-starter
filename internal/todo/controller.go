package todo

import (
	"errors"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	repo Repository
}

type LastTodoRequest struct{}

type LastTodoResponse struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

func NewController(repo Repository) *Controller {
	return &Controller{repo}
}

func (cc *Controller) LastTodo(c echo.Context) error {
	var req LastTodoRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	t, err := cc.repo.FindLast(c.Request().Context())
	if err != nil {
		return mapError(err)
	}

	return c.JSON(http.StatusOK, LastTodoResponse{
		ID:          t.ID,
		Description: t.Description,
		DueDate:     t.DueDate.Format(time.RFC3339),
	})
}

func mapError(err error) error {
	switch {
	case errors.Is(err, ErrTodoNotFound):
		return echo.NewHTTPError(http.StatusNotFound)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
}
