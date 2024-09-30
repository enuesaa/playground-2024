package files

import (
	"path/filepath"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)
	filename := cc.Param("filename")

	path := filepath.Join(cc.Repos.Config.DocsPath, filename)
	fbytes, err := cc.Repos.Fs.Read(path)
	if err != nil {
		return err
	}
	content := string(fbytes)

	res := Detail{
		Content: content,
	}
	return cc.JSON(200, res)
}
