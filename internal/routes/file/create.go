package file

import (
	"strings"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)

	var reqbody CreateRequestBody
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	reader := strings.NewReader(reqbody.Content)
	if err := cc.Repos.Fs.Create(reqbody.Path, reader); err != nil {
		return err
	}

	res := Creation{
		Ok: true,
	}
	return cc.JSON(200, res)
}
