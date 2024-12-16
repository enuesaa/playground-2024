package callerfx

import "go.uber.org/fx"

var Module = fx.Module(
	"callerfx",
	fx.Provide(New),
)
