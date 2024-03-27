package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/", func (c echo.Context) error {
		res, err := http.Get("https://example.com")
		if err != nil {
			return err
		}
		defer res.Body.Close()
		return c.String(200, "aaa")
	})

	c := jaegertracing.New(app, nil)
	defer c.Close()
	
	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
