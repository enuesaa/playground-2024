package filetree

import (
	"bytes"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)

	buffer := bytes.NewBuffer([]byte{})
	if err := cc.Repos.Cmd.Exec("tree", buffer); err != nil {
		return err
	}

	res := Item{
		Tree: buffer.String(),
	}
	return cc.WithData(res)
}
