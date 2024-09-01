package serve

import (
	"github.com/labstack/echo/v4"
)

type GetUserResponse struct {
	Name string `json:"name"`
}

func (ctl *ServeCtl) handleGetUser(c echo.Context) error {
	res := GetUserResponse{
		Name: "a",
	}
	return c.JSON(200, res)
}
