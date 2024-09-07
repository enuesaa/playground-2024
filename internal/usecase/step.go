package usecase

import (
	"fmt"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
)

func WriteStep(repos repository.Repos) error {
	// TODO
	// .console という文字列を打ったらプロンプトが立ち上がる
	// .capture という文字列を打ったら menu が立ち上がる

	texts := []string{
		"# ",
		"",
		"## ",
		"",
		"",
	}

	for _, text := range texts {
		if _, err := repos.Log.Ask(">", text); err != nil {
			return err
		}
	}

	fmt.Println("end")
	input := textinput.New("")
	input.Validate = nil
	input.AutoComplete = textinput.AutoCompleteFromSlice([]string{"```console"})

	text, _ := input.RunPrompt()
	fmt.Println(text)

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
