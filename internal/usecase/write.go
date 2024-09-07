package usecase

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit/textinput"
)

func Write(repos repository.Repos, filename string) error {
	texts := []string{}

	for {
		input := textinput.New(">")
		input.InitialValue = ""
		input.Validate = nil
		input.Template = `
		{{- A }}
		{{- Bold .Prompt }} {{ .Input -}}
		{{- if .ValidationError }} {{ Foreground "1" (Bold "✘") }}
		{{- end -}}
		`
		input.ExtendedTemplateFuncs = template.FuncMap{
			"A": func () string {
				// fmt.Println("####")
				return "####"
			},
		}
		text, err := input.RunPrompt()
		if err != nil {
			return err
		}
		if text == ".console" {
			result, err := Prompt(repos)
			if err != nil {
				return err
			}
			texts = append(texts, result)
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
