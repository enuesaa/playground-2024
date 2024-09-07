package usecase

import (
	"fmt"
	"strings"

	"github.com/enuesaa/codetrailer/internal/repository"
)

func Write(repos repository.Repos, filename string) error {
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

	body := strings.Join(texts, "\n")
	bodyreader := strings.NewReader(body)

	path := fmt.Sprintf(".codetrailer/%s.md", filename)
	if err := repos.Fs.Create(path, bodyreader); err != nil {
		return err
	}

	return nil
}
