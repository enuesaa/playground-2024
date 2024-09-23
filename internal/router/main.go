package router

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/router/middleware"
	apiDocs "github.com/enuesaa/codetrailer/internal/routes/docs"
	apiFile "github.com/enuesaa/codetrailer/internal/routes/file"
	apiFiles "github.com/enuesaa/codetrailer/internal/routes/files"
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

	api.GET("/docs", apiDocs.List)
	api.GET("/docs/:name", apiDocs.View)
	api.POST("/docs", apiDocs.Create)
	api.PUT("/docs/:name", apiDocs.Update)

	api.GET("/files", apiFiles.List)
	// api.GET("/file", apiFile.Create)
	api.POST("/file", apiFile.Create)
	api.GET("/filetree", apiFiletree.View)

	api.POST("/prompt", apiPrompt.Create)
	// api.POST("/prompt/{id}/command", apiFiles.List)

	app.Any("/*", ui.Serve)

	return app
}
