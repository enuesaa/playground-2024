package routes

import (
	"strings"

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
	for _, f := range files {
		isDir, err := cc.Repos.Fs.IsDir(f)
		if err != nil {
			isDir = false
		}

		splitted := strings.Split(f, "/")
		filename := splitted[len(splitted)-1]
		items = append(items, Item{
			Path:     f,
			Filename: filename,
			IsDir:    isDir,
		})
	}

	return cc.JSON(200, items)
}
