package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

var Cors = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"http://localhost:3001"},
})
