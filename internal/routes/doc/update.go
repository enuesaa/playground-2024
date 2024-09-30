package doc

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	cc := ctx.Use(c)

	var reqbody UpdateRequestBody
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	docDir := filepath.Dir(cc.Repos.Config.DocPath)

	readme := ""
	for i, slide := range reqbody.Slides {
		readme += slide.Content + "\n\n---\n\n"
		drawingPath := filepath.Join(docDir, fmt.Sprintf("%d.svg", i))
		if err := cc.Repos.Fs.Create(drawingPath, strings.NewReader(slide.Drawing)); err != nil {
			continue
		}
	}

	reader := strings.NewReader(readme)
	if err := cc.Repos.Fs.Create(cc.Repos.Config.DocPath, reader); err != nil {
		return err
	}

	res := Creation{
		Ok: true,
	}
	return cc.JSON(200, res)
}
