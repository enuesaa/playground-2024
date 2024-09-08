package usecase

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit"
)

func Prompt(repos repository.Repos) (string, error) {
	var result bytes.Buffer
	output := io.MultiWriter(os.Stdout, &result)

	for {
		args, err := repos.Prompt.Render("", func(s string) (string, bool) {
			if strings.HasPrefix(s, "@") {
				return s, false
			}
			return fmt.Sprintf("$ %s", s), true
		})
		if err != nil {
			if errors.Is(err, promptkit.ErrAborted) {
				return result.String(), nil
			}
			return "", err
		}
		if strings.HasPrefix(args, "@") {
			return result.String(), nil
		}

		executed := fmt.Sprintf("$ %s\n", args)
		if _, err := result.Write([]byte(executed)); err != nil {
			return result.String(), err
		}

		if err := repos.Cmd.Exec(args, output); err != nil {
			return result.String(), nil
		}
		if _, err := output.Write([]byte("\n")); err != nil {
			return result.String(), err
		}
	}
}
