package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

func setupTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, err
	}

	resources, err := newTraceResoure(ctx)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)
	otel.SetTracerProvider(tp)

	return tp, nil
}

func newTraceResoure(ctx context.Context) (*resource.Resource, error) {
	return resource.New(
		ctx,
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceName("aa"),
		),
	)
}
