package files

import (
	"path/filepath"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	cc := ctx.Use(c)
	filename := cc.Param("filename")

	path := filepath.Join(cc.Repos.Config.DocsPath, filename)
	if err := cc.Repos.Fs.Delete(path); err != nil {
		return err
	}

	res := Creation{
		Ok: true,
	}
	return cc.JSON(200, res)
}
