package doc

import (
	"strings"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)

	fbytes, err := cc.Repos.Fs.Read(cc.Repos.Config.DocPath)
	if err != nil {
		return err
	}
	content := string(fbytes)
	slides := strings.Split(content, "\n---\n")

	res := Detail{
		Path: cc.Repos.Config.DocPath,
		Slides: slides,
	}
	return cc.JSON(200, res)
}
