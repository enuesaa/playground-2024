package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	// this is normal lambda handler
	// lambda.Start(Handler)

	// using lambda web adapter
	app := echo.New()
	app.GET("/hello", func (c echo.Context) error {
		return c.String(200, "hello")
	})
	app.GET("/hello/hello", func (c echo.Context) error {
		return c.String(200, "hello hello")
	})

	// lambda web adapter のデフォルトポートは 8080
	// see https://github.com/awslabs/aws-lambda-web-adapter/blob/35dc4cc58e5a52042e1bc31058e3f8abecc9e4fd/src/lib.rs#L91
	if err := app.Start(":8080"); err != nil {
		panic(err)
	}
}
