package usecase

import (
	"path/filepath"
	"strings"

	"github.com/enuesaa/codetrailer/internal/repository"
)

func Read(repos repository.Repos, path string) []string {
	readme := filepath.Join(path, "README.md")

	content, err := repos.Fs.Read(readme)
	if err != nil {
		return make([]string, 0)
	}

	return strings.Split(string(content), "\n")
}
