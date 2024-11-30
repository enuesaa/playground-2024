package router

import (
	"net/http"

	"github.com/enuesaa/sample-books-api/internal/router/middleware"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	app := echo.New()

	// middleware
	app.Use(middleware.Cors())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return app
}
