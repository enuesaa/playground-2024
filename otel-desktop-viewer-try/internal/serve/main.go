package serve

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func New() ServeCtl {
	return ServeCtl{}
}

type ServeCtl struct{
	tracerProvider *sdktrace.TracerProvider
	tracer trace.Tracer
}

func (ctl *ServeCtl) Serve() error {
	if err := ctl.setupTracer(); err != nil {
		return err
	}

	// app
	app := echo.New()

	// middleware
	app.Use(otelecho.Middleware("notes"))
	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				return err
			}

			ctx := c.Request().Context()
			span := trace.SpanFromContext(ctx)
			if span != nil {
				status := c.Response().Status
				span.SetAttributes(attribute.Int("http.status_code", status))
				if status < 400 {
					span.SetStatus(codes.Ok, "Ok")
				} else {
					span.SetStatus(codes.Error, "Error")
				}
			}
			return nil
		}
	})
	app.Use(ctl.handleUsefulMiddleware)

	// routing
	app.GET("/users/:id", ctl.handleGetUser)

	return app.Start(":3000")
}

func (ctl *ServeCtl) Shutdown() {
	if ctl.tracerProvider != nil {
		if err := ctl.tracerProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error: %s\n", err.Error())
		}
	}
}
