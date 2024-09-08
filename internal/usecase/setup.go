package usecase

import (
	"github.com/enuesaa/codetrailer/internal/repository"
)

func Setup(repos repository.Repos, path string) error {
	if repos.Fs.IsExist(path) {
		return nil
	}
	return repos.Fs.CreateDir(path)
}
