package file

import (
	"fmt"
	"strings"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)

	var reqbody CreateBody
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	fmt.Println(reqbody.Path)
	fmt.Println(reqbody.Content)

	reader := strings.NewReader(reqbody.Content)
	if err := cc.Repos.Fs.Create(reqbody.Path, reader); err != nil {
		return err
	}

	res := Created{
		Ok: true,
	}
	return cc.WithData(res)
}
