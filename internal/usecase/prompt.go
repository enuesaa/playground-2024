package usecase

import (
	"errors"
	"os"
	"os/exec"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit"
)

func Prompt(repos repository.Repos) error {
	for {
		args, err := repos.Log.Ask(">", "")
		if err != nil {
			if errors.Is(err, promptkit.ErrAborted) {
				return nil
			}
			return err
		}
		if args == "q" {
			break
		}
		cmd := exec.Command("bash", "-c", args)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
