package middleware

import (
	"github.com/labstack/echo/v4"
)

type ApiResponse struct {
	Data interface{} `json:"data"`
}

func HandleData(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			return err
		}

		// check that reponse body is already set.
		if c.Response().Committed {
			return nil
		}

		data := c.Get("data")
		if data == nil {
			data = struct{}{}
		}

		res := ApiResponse{
			Data: data,
		}
		return c.JSON(200, res)
	}
}
