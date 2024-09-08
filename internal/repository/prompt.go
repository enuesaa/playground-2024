package repository

import (
	"github.com/erikgeiser/promptkit/textinput"
)

type PromptRepositoryInterface interface {
	Ask(message string, defaultValue string) (string, error)
}

type PromptRepository struct {}

func (repo *PromptRepository) Ask(message string, defaultValue string) (string, error) {
	input := textinput.New(message)
	input.InitialValue = defaultValue
	input.Validate = nil
	input.Template = `
	{{- Bold .Prompt }} {{ .Input -}}
	{{- if .ValidationError }} {{ Foreground "1" (Bold "✘") }}
	{{- end -}}
	`

	return input.RunPrompt()
}
