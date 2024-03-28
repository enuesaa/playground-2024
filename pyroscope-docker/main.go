package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/grafana/pyroscope-go"
)

func main() {
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "sample.app",
		ServerAddress:   "http://localhost:4040",
	})

	app := echo.New()
	app.GET("/", func (c echo.Context) error {
		res, err := http.Get("https://example.com")
		if err != nil {
			return err
		}
		defer res.Body.Close()
		return c.String(200, "aaa")
	})
	
	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
