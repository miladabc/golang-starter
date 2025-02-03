package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/miladabc/golang-starter/internal/http/httptest"
	"github.com/stretchr/testify/require"
)

func TestControllerLastTodo(t *testing.T) {
	t.Parallel()

	repo := NewMockRepo(MockRepoOpt{Err: nil})
	con := NewController(repo)
	c, res := httptest.NewEchoContext(t)

	err := con.LastTodo(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.Code)

	expectedTodo := LastTodoResponse{
		ID:          1,
		Description: "mock todo",
		DueDate:     "2025-01-01T12:00:00Z",
	}

	var actualTodo LastTodoResponse

	err = json.Unmarshal(res.Body.Bytes(), &actualTodo)
	require.NoError(t, err)
	require.Equal(t, expectedTodo, actualTodo)
}

func TestControllerLastTodo_NotFound(t *testing.T) {
	t.Parallel()

	repo := NewMockRepo(MockRepoOpt{Err: ErrTodoNotFound})
	con := NewController(repo)
	c, _ := httptest.NewEchoContext(t)

	err := con.LastTodo(c)

	var hErr *echo.HTTPError

	require.ErrorAs(t, err, &hErr)
	require.Equal(t, http.StatusNotFound, hErr.Code)
}

func TestControllerLastTodo_Err(t *testing.T) {
	t.Parallel()

	dErr := fmt.Errorf("db error")
	repo := NewMockRepo(MockRepoOpt{Err: dErr})
	con := NewController(repo)
	c, _ := httptest.NewEchoContext(t)

	err := con.LastTodo(c)
	require.ErrorIs(t, err, dErr)

	var hErr *echo.HTTPError

	require.ErrorAs(t, err, &hErr)
	require.Equal(t, http.StatusInternalServerError, hErr.Code)
}
