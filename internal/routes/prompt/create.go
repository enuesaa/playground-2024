package prompt

import (
	"bytes"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)

	var reqbody CreateBody
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	buffer := bytes.NewBuffer([]byte{})
	if err := cc.Repos.Cmd.Exec(reqbody.Command, buffer); err != nil {
		return err
	}

	res := Created{
		Output: buffer.String(),
	}
	return cc.WithData(res)
}
