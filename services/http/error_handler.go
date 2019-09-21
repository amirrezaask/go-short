package http

import (
	"github.com/labstack/echo"
	"go-short/services/responses"
	"net/http"
)

func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if err := c.JSON(code, responses.Error(err.Error())); err != nil {
		panic(err)
	}

	c.Logger().Error(err)
}
