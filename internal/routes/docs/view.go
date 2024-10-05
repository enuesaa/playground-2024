package docs

import (
	"fmt"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)
	name := cc.Param("name")

	docPath := fmt.Sprintf("%s/%s/README.md", cc.Repos.Config.DocsPath, name)
	fbytes, err := cc.Repos.Fs.Read(docPath)
	if err != nil {
		return err
	}
	content := string(fbytes)

	res := Detail{
		Path:    fmt.Sprintf("%s/%s", cc.Repos.Config.DocsPath, name),
		DirName: name,
		Content: content,
	}
	return cc.JSON(200, res)
}
