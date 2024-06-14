package main

import (
	"log"

	"github.com/evanw/esbuild/pkg/api"
	"rogchap.com/v8go"
)

func main() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"index.cjs"},
		Bundle: true,
		Write: false,
		Format: api.FormatCommonJS,
		Platform: api.PlatformNode,
		Target: api.ES2015,
		// Banner: map[string]string{
		// 	"js": `import { createRequire } from "module"; import url from "url"; const require = createRequire(import.meta.url); const __filename = url.fileURLToPath(import.meta.url); const __dirname = url.fileURLToPath(new URL(".", import.meta.url));`,
		// },
	})
	log.Println(result)
	app := string(result.OutputFiles[0].Contents)

	ctx := v8go.NewContext()
	_, err := ctx.RunScript(app, "main.js")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
