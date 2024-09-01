package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"  
)

var tracer = otel.Tracer("echo-server")

func main() {
	tp, err := initTracer()
	if err != nil {
		panic(err)
	}
	defer tp.Shutdown(context.Background())

	r := echo.New()
	r.Use(otelecho.Middleware("hello"))
	r.Use(handleUsefulMiddleware)

	// routing
	r.GET("/users/:id", handleGetUser)

	if err := r.Start(":3000"); err != nil {
		panic(err)
	}
}

func initTracer() (*sdktrace.TracerProvider, error) {
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func handleUsefulMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// この context を引き回さないと span がまとまらない
		_, span := tracer.Start(c.Request().Context(), "useful-middleware")
		// _, span := tracer.Start(ctx, "getUser", oteltrace.WithAttributes(attribute.String("id", id)))
		defer span.End()
		return next(c)
	}
}

type GetUserResponse struct {
	Name string `json:"name"`
}
func handleGetUser(c echo.Context) error {
	res := GetUserResponse{
		Name: "a",
	}
	return c.JSON(http.StatusOK, res)
}
