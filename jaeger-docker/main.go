package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"

	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	// "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
    _ "go.opentelemetry.io/otel/semconv/v1.9.0"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var tracer = otel.Tracer("echo-server")

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer tp.Shutdown(context.Background())

	app := echo.New()
	app.Use(otelecho.Middleware("a"))

	app.GET("/users/:id", func(c echo.Context) error {
		return c.String(200, "aaa")
	})
	if err = app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func initTracer() (*sdktrace.TracerProvider, error) {
	// client := otlptracehttp.NewClient(
	// 	otlptracehttp.WithEndpoint("10.20.30.40:4318"),
	// 	otlptracehttp.WithInsecure(),
	// 	otlptracehttp.WithCompression(otlptracehttp.NoCompression),
	// )
	// exporter, err := otlptrace.New(context.Background(), client)
	// if err != nil {
	// 	return nil, err
	// }

	endpoint := "http://localhost:14268/api/traces"
	exp, err := jaeger.NewRawExporter(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}
