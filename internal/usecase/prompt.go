package usecase

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit"
)

func Prompt(repos repository.Repos) (string, error) {
	var result bytes.Buffer
	outputWriter := io.MultiWriter(os.Stdout, &result)

	for {
		args, err := repos.Log.Ask("$", "")
		if err != nil {
			if errors.Is(err, promptkit.ErrAborted) {
				return "", nil
			}
			return "", nil
		}
		if args == "@exit" {
			return result.String(), nil
		}
		if _, err := result.Write([]byte(args)); err != nil {
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
