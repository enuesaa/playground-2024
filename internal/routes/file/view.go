package file

import (
	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)
	path := cc.QueryParam("path")

	fbytes, err := cc.Repos.Fs.Read(path)
	if err != nil {
		return err
	}
	content := string(fbytes)

	res := Detail{
		Path: path,
		Content: content,
	}
	return cc.JSON(200, res)
}
