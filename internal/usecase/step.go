package usecase

import (
	"fmt"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit/selection"
)

func WriteStep(repos repository.Repos) error {
	sp := selection.New("Which", []string{"Title", "Description", "Sh", "Capture"})
	sp.Filter = nil
	choice, err := sp.RunPrompt()
	if err != nil {
		return err
	}

	if choice == "Title" {
		title, err := repos.Log.Ask("title>", "")
		if err != nil {
			return err
		}
		fmt.Println(title)
	}

	if choice == "Description" {
		title, err := repos.Log.Ask("description>", "")
		if err != nil {
			return err
		}
		fmt.Println(title)
	}

	if choice == "Sh" {
		if err := Prompt(repos); err != nil {
			return err
		}
	}

	if choice == "Capture" {
		if err := LaunchMenu(); err != nil {
			return err
		}
	}

	return nil
}
