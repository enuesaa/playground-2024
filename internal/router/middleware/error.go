package middleware

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

type AppErr struct {
	error
	Errors []AppErrItem `json:"errors"`
}
type AppErrItem struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

func NewAppErr() AppErr {
	return AppErr{
		Errors: make([]AppErrItem, 0),
	}
}

func HandleError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		if errors.As(err, &AppErr{}) {
			return c.JSON(422, err)
		}
		// like 404
		if _, ok := err.(*echo.HTTPError); ok {
			return err
		}
		fmt.Printf("Error: %s", err)

		apperr := AppErr{
			Errors: []AppErrItem{
				{
					Message: "Internal Server Error",
				},
			},
		}
		return c.JSON(500, apperr)
	}
}
