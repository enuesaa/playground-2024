package usecase

import (
	"github.com/enuesaa/codetrailer/internal/repository"
)

func CreateRegistry(repos repository.Repos) error {
	if repos.Fs.IsExist(".codetrailer") {
		return nil
	}
	if err := repos.Fs.CreateDir(".codetrailer"); err != nil {
		return err
	}

	return nil
}
