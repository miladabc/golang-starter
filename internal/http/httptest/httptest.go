package httptest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/miladabc/golang-starter/pkg/validator"
)

func NewEchoContext(t *testing.T) (echo.Context, *httptest.ResponseRecorder) {
	t.Helper()

	r := echo.New()
	r.Debug = true
	r.Validator = validator.NewEchoValidator()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	return r.NewContext(req, res), res
}
