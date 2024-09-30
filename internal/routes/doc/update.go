package doc

import (
	"strings"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	cc := ctx.Use(c)

	var reqbody UpdateRequestBody
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	content := strings.Join(reqbody.Slides, "\n---\n")
	reader := strings.NewReader(content)
	if err := cc.Repos.Fs.Create(cc.Repos.Config.DocsPath, reader); err != nil {
		return err
	}

	res := Updated{
		Ok: true,
	}
	return cc.JSON(200, res)
}
