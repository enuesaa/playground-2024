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

	// AppCommand
	//   @ で始まるのが AppCommand
	//   @ で始まったら行頭の > を消して @ を表示する
	//   - @console
	//   - @exit

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
		if text == "@console" {
			result, err := Prompt(repos)
			if err != nil {
				return err
			}
			texts = append(texts, result)
		} else if text == "@exit" {
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
