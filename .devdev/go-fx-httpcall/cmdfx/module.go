package cmdfx

import "go.uber.org/fx"

var Module = fx.Module(
	"cmdfx",
	fx.Provide(New),
)
