package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Recover() echo.MiddlewareFunc {
	handle := middleware.Recover()

	return handle
}
