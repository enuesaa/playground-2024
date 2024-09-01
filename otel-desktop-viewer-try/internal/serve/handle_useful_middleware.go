package serve

import (
	"github.com/labstack/echo/v4"

	"go.opentelemetry.io/otel/attribute"
)

func (ctl *ServeCtl) handleUsefulMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// この context を引き回さないと span がまとまらない

		ctx := c.Request().Context()

		_, span := ctl.tracer.Start(ctx, "useful-middleware")
		span.SetAttributes(attribute.String("hey", "hello"))
		defer span.End()

		return next(c)
	}
}
