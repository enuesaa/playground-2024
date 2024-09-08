package usecase

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/creack/pty"
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit"
)

func Prompt(repos repository.Repos) (string, error) {
	var result bytes.Buffer
	outputWriter := io.MultiWriter(os.Stdout, &result)

	for {
		args, err := repos.Prompt.Render("", func(s string) (string, bool) {
			if strings.HasPrefix(s, "@") {
				return s, false
			}
			return fmt.Sprintf("$ %s", s), true
		})
		if strings.HasPrefix(args, "@") {
			return result.String(), nil
		}
		if err != nil {
			if errors.Is(err, promptkit.ErrAborted) {
				return result.String(), nil
			}
			return "", err
		}
		executed := fmt.Sprintf("$ %s\n", args)
		if _, err := result.Write([]byte(executed)); err != nil {
			return result.String(), err
		}

		cmd := exec.Command("bash", "-c", args)

		pf, err := pty.Start(cmd)
		if err != nil {
			return result.String(), nil
		}
		defer pf.Close()

		io.Copy(outputWriter, pf)
	}
}
