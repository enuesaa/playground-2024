package clifx

import "go.uber.org/fx"

var Module = fx.Module(
	"clifx",
	fx.Provide(New),
)
