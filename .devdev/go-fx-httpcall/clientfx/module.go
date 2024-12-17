package clientfx

import "go.uber.org/fx"

var Module = fx.Module(
	"clientfx",
	fx.Provide(New),
)
