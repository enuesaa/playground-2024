package files

import (
	"path/filepath"
	"strings"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	cc := ctx.Use(c)
	filename := cc.Param("filename")

	var reqbody UpdateRequestBody
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	path := filepath.Join(cc.Repos.Config.DocsPath, filename)
	reader := strings.NewReader(reqbody.Content)
	if err := cc.Repos.Fs.Create(path, reader); err != nil {
		return err
	}

	res := Creation{
		Ok: true,
	}
	return cc.JSON(200, res)
}
