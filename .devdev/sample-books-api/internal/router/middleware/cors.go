package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Cors() echo.MiddlewareFunc {
	handle := middleware.CORS()

	return handle
}
