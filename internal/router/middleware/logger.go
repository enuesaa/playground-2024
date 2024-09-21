package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

var Logger = middleware.LoggerWithConfig(middleware.LoggerConfig{
	Format: "method=${method}, uri=${uri}, status=${status}\n",
})
