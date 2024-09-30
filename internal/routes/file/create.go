package file

import (
	"path/filepath"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)
	filename := cc.Param("filename")
	basepath := filepath.Dir(cc.Repos.Config.DocPath)
	path := filepath.Join(basepath, filename)

	file, err := cc.FormFile("file")
	if err != nil {
		return err
	}

	f, err := file.Open()
	if err != nil {
		return err
	}
	if err := cc.Repos.Fs.Create(path, f); err != nil {
		return err
	}

	res := Creation{
		Ok: true,
	}
	return cc.JSON(200, res)
}
