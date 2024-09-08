package repository

import (
	"github.com/playwright-community/playwright-go"
)

type PwRepositoryInterface interface {
	Install() error
}

type PwRepository struct{}

func (repo *PwRepository) Install() error {
	options := playwright.RunOptions{
		Browsers: []string{"chromium"},
	}
	return playwright.Install(&options)
}
