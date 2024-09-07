package usecase

import (
	"fmt"

	"github.com/enuesaa/codetrailer/internal/repository"
)

func Write(repos repository.Repos) error {
	texts := []string{}

	for {
		text, err := repos.Log.Ask(">", "#")
		if err != nil {
			return err
		}
		if text == ".console" {
			if err := Prompt(repos); err != nil {
				return err
			}
		} else if text == ".exit" {
			fmt.Println(texts)
			break
		} else {
			texts = append(texts, text)
		}
	}

	return nil
}
