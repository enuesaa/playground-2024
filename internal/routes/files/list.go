package routes

import (
	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	cc := ctx.Use(c)
	path := cc.QueryParam("path")

	items := []Item{}

	files, err := cc.Repos.Fs.ListFiles(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		isDir, err := cc.Repos.Fs.IsDir(file)
		if err != nil {
			isDir = false
		}
		items = append(items, Item{
			Name: file,
			IsDir: isDir,
		})
	}

	return cc.WithData(files)
}
