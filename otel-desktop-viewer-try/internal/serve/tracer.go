package serve

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

func (ctl *ServeCtl) setupTracer() error {
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		return err
	}

	// root span に service name をセットする
	rsource, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("my app"),
		),
	)
	if err != nil {
		return err
	}

	ctl.tracerProvider = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(rsource),
	)
	otel.SetTracerProvider(ctl.tracerProvider)

	ctl.tracer = ctl.tracerProvider.Tracer("notes")

	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	otel.SetTextMapPropagator(propagator)

	return nil
}
