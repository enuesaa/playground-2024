package doc

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)

	docDir := filepath.Dir(cc.Repos.Config.DocPath)

	fbytes, err := cc.Repos.Fs.Read(cc.Repos.Config.DocPath)
	if err != nil {
		return err
	}
	readme := string(fbytes)

	slides := []Slide{}
	for i, content := range strings.Split(readme, "\n\n---\n\n") {
		drawingPath := filepath.Join(docDir, fmt.Sprintf("%d.svg", i))
		fbytes, err := cc.Repos.Fs.Read(drawingPath)
		if err != nil {
			slides = append(slides, Slide{
				Content: content,
				Drawing: "",
			})
			continue
		}
		slides = append(slides, Slide{
			Content: content,
			Drawing: string(fbytes),
		})
	}

	res := Detail{
		Slides: slides,
	}
	return cc.JSON(200, res)
}
