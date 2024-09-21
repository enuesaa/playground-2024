package router

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/router/middleware"
	routes "github.com/enuesaa/codetrailer/internal/routes"
	"github.com/enuesaa/codetrailer/ui"

	"github.com/labstack/echo/v4"
)

func New(repos repository.Repos) *echo.Echo {
	app := echo.New()
	app.Use(middleware.BindCtx(repos))
	app.Use(middleware.Logger)
	app.Use(middleware.Cors)

	api := app.Group("/api")
	api.Use(middleware.HandleData)
	api.Use(middleware.HandleError)

	api.GET("/", routes.List)

	app.Any("/*", ui.Serve)

	return app
}
