package ui

import (
	"embed"
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/enuesaa/codetrailer/internal/router/ctx"
)

//go:generate pnpm install
//go:generate pnpm build

//go:embed all:dist/*
var dist embed.FS

func Serve(c echo.Context) error {
	cc := ctx.Use(c)

	path := cc.Request().URL.Path // like `/index.html`
	path = strings.TrimPrefix(path, "/") // like `index.html`

	distpath := fmt.Sprintf("dist/%s", path)
	fileExt := filepath.Ext(distpath)
	if fileExt == "" {
		distpath = "dist/index.html"
	}

	f, err := dist.ReadFile(distpath)
	if err != nil {
		// read files in current path
		assetdir := filepath.Dir(cc.Repos.Config.DocPath)
		assetpath := filepath.Join(assetdir, path)
		f, err = cc.Repos.Fs.Read(assetpath)
		if err != nil {
			return err
		}
	}
	mimeType := mime.TypeByExtension(fileExt)

	return c.Blob(200, mimeType, f)
}
