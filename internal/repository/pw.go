package repository

import (
	"github.com/playwright-community/playwright-go"
)

type PwRepositoryInterface interface {
	Install() error
}

type PwRepository struct{}

func (repo *PwRepository) Install() error {
	err := playwright.Install(&playwright.RunOptions{
		Browsers: []string{"chromium"},
	})
	if err != nil {
		return err
	}
	return nil
}
