package web

import (
	"embed"
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

//go:generate pnpm install
//go:generate pnpm build

//go:embed all:dist/*
var dist embed.FS

var Serve = ServeDist

func ServeDist(c *fiber.Ctx) error {
	path := c.Path() // like `/`
	path = fmt.Sprintf("dist%s", path)
	if strings.HasSuffix(path, "/") {
		path += "index.html"
	}

	f, err := dist.ReadFile(path)
	if err != nil {
		return err
	}
	fileExt := filepath.Ext(path)
	mimeType := mime.TypeByExtension(fileExt)
	c.Set(fiber.HeaderContentType, mimeType)

	return c.SendString(string(f))
}
