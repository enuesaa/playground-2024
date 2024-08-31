package usecase

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit"
)

func Prompt(repos repository.Repos) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	stdoutWriter := io.MultiWriter(os.Stdout, &stdout)
	stderrWriter := io.MultiWriter(os.Stderr, &stderr)

	for {
		args, err := repos.Log.Ask(">", "")
		if err != nil {
			if errors.Is(err, promptkit.ErrAborted) {
				return nil
			}
			return err
		}
		if args == "q" {
			fmt.Printf("captured: \n%s\n", stdout.String())
			fmt.Printf("captured: \n%s\n", stderr.String())
			break
		}
		cmd := exec.Command("bash", "-c", args)
		cmd.Stdout = stdoutWriter
		cmd.Stderr = stderrWriter

		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
