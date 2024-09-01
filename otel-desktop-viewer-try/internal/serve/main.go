package serve

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
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
