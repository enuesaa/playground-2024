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

func Prompt(repos repository.Repos) error {
	var result bytes.Buffer
	outputWriter := io.MultiWriter(os.Stdout, &result)

	for {
		args, err := repos.Log.Ask("console>", "")
		if err != nil {
			if errors.Is(err, promptkit.ErrAborted) {
				return nil
			}
			return err
		}
		if args == "q" {
			return repos.Fs.Create("out.txt", &result)
		}

		cmd := exec.Command("bash", "-c", args)

		pf, err := pty.Start(cmd)
		if err != nil {
			return err
		}
		defer pf.Close()

		io.Copy(outputWriter, pf)
	}
}
