package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, "hello new World from Japan")
	}
}
