package ctx

import (
	"github.com/labstack/echo/v4"
)

func Use(c echo.Context) *Context {
	return c.(*Context)
}
