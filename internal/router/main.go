package router

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/router/middleware"
	apiDoc "github.com/enuesaa/codetrailer/internal/routes/doc"
	apiFile "github.com/enuesaa/codetrailer/internal/routes/file"
	apiFiletree "github.com/enuesaa/codetrailer/internal/routes/filetree"
	apiPrompt "github.com/enuesaa/codetrailer/internal/routes/prompt"
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

	api.GET("/doc", apiDoc.View)
	api.PUT("/doc", apiDoc.Update)
	api.POST("/file/:filename", apiFile.Create)

	api.GET("/filetree", apiFiletree.View)

	api.POST("/prompt", apiPrompt.Create)

	app.Any("/*", ui.Serve)

	return app
}
