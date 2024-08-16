package main

import (
	_ "gf-try/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-try/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
