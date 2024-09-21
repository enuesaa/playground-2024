package ctx

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context

	Repos repository.Repos
}

func (cc *Context) WithData(data interface{}) error {
	cc.Set("data", data)

	return nil
}
