package usecase

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit"
	"github.com/getlantern/systray"
)

// AppCommand
//   @ で始まるのが AppCommand
//   @ で始まったら行頭の > を消して @ を表示する
//   - @sh
//   - @
//   - @cat
func Write(repos repository.Repos, path string) error {
	texts := []string{}
	existing := Read(repos, path)

	i := 0
	for {
		defaultValue := ""
		if len(existing) > i {
			defaultValue = existing[i]
		}
		text, err := repos.Prompt.Render(defaultValue, func(s string) (string, bool) {
			if strings.HasPrefix(s, "@") {
				return s, false
			}
			return fmt.Sprintf("> %s", s), true
		})
		if err != nil {
			if errors.Is(err, promptkit.ErrAborted) {
				break
			}
			return err
		}
		if text == "@sh" {
			result, err := Prompt(repos)
			if err != nil {
				return err
			}
			texts = append(texts, "```console")
			texts = append(texts, result)
			texts = append(texts, "```")
		} else if text == "@" {
			break
		} else {
			texts = append(texts, text)
		}
		i++
	}

	body := strings.Join(texts, "\n")
	bodyreader := strings.NewReader(body)
	readme := filepath.Join(path, "README.md")

	if err := repos.Fs.Create(readme, bodyreader); err != nil {
		return err
	}
	systray.Quit()

	return nil
}
