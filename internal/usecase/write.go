package usecase

import (
	"path/filepath"
	"strings"
	"text/template"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit/textinput"
)

// AppCommand
//   @ で始まるのが AppCommand
//   @ で始まったら行頭の > を消して @ を表示する
//   - @sh
//   - @q
//   - @cat
func Write(repos repository.Repos, path string) error {
	texts := []string{}

	for {
		input := textinput.New(">")
		input.InitialValue = ""
		input.Validate = nil
		input.Template = `
		{{- if IsAppCommand .Input }}
		{{- .Input -}}
		{{- else -}}
		{{- Bold .Prompt }} {{ .Input -}}
		{{- end -}}
		{{- if .ValidationError }} {{ Foreground "1" (Bold "✘") }}
		{{- end -}}
		`
		input.ResultTemplate = `
		{{- if IsAppCommand .FinalValue }}
		{{- else -}}
		{{- print .Prompt " " (Foreground "32"  (Mask .FinalValue)) "\n" -}}
		{{- end -}}
		`
		input.ExtendedTemplateFuncs = template.FuncMap{
			"IsAppCommand": func (s string) bool {
				return strings.HasPrefix(s, "@")
			},
		}
		text, err := input.RunPrompt()
		if err != nil {
			return err
		}
		if text == "@sh" {
			result, err := Prompt(repos)
			if err != nil {
				return err
			}
			texts = append(texts, result)
		} else if text == "@q" {
			break
		} else {
			texts = append(texts, text)
		}
	}

	body := strings.Join(texts, "\n")
	bodyreader := strings.NewReader(body)
	readme := filepath.Join(path, "README.md")

	return repos.Fs.Create(readme, bodyreader)
}
