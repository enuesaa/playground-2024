package docs

import (
	"strings"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	cc := ctx.Use(c)
	items := []Item{}

	docsPath := cc.Repos.Config.DocsPath
	files, err := cc.Repos.Fs.ListFiles(docsPath)
	if err != nil {
		return err
	}
	for _, f := range files {
		isDir, err := cc.Repos.Fs.IsDir(f)
		if err != nil {
			continue
		}
		// In this app, `docs` means README.md inside the subjective directory like `docs/aaa/README.md`
		if !isDir {
			continue
		}

		splitted := strings.Split(f, "/")
		dirname := splitted[len(splitted)-1]
		items = append(items, Item{
			Path:    f,
			DirName: dirname,
		})
	}

	return cc.JSON(200, items)
}
