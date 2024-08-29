package main

import (
	"log"

	"github.com/labstack/echo/v4"

	apmecho "go.elastic.co/apm/module/apmechov4/v2"
	"go.elastic.co/apm/v2"
)

func main() {
	e := echo.New()
	e.Use(apmecho.Middleware())

	e.GET("/hello/:name", func(c echo.Context) error {
		span, _ := apm.StartSpan(c.Request().Context(), "work", "custom")
		defer span.End()
		return nil
	})

	if err := e.Start(":3000"); err != nil {
		log.Panic(err.Error())
	}
}
