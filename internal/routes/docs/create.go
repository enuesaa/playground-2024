package docs

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)

	var reqbody CreateBody
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	path := filepath.Join(cc.Repos.Config.DocsPath, reqbody.DirName)
	if err := cc.Repos.Fs.CreateDir(path); err != nil {
		return err
	}

	readmePath := filepath.Join(cc.Repos.Config.DocsPath, reqbody.DirName, "README.md")
	readme := fmt.Sprintf("# %s\n## 概要", reqbody.DirName)

	reader := strings.NewReader(readme)
	if err := cc.Repos.Fs.Create(readmePath, reader); err != nil {
		return err
	}

	res := Created{
		Ok: true,
	}
	return cc.WithData(res)
}