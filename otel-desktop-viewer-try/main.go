package main

import (
	"gihtub.com/enuesaa/playground-2024/otel-desktop-viewer-try/internal/serve"
	"log"
)

func main() {
	serveCtl := serve.New()
	defer serveCtl.Shutdown()

	if err := serveCtl.Serve(); err != nil {
		log.Panicf("Error: %s\n", err.Error())
	}
}
