package usecase

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/enuesaa/codetrailer/internal/repository"
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
		text, err := repos.Prompt.Render("", func(s string) (string, bool) {
			if strings.HasPrefix(s, "@") {
				return s, false
			}
			return fmt.Sprintf("> %s", s), true
		})
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
