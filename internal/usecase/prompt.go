package usecase

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit"
)

func Prompt(repos repository.Repos) error {
	var stdout bytes.Buffer
	stdoutWriter := io.MultiWriter(os.Stdout, &stdout)

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
			break
		}
		cmd := exec.Command("bash", "-c", args)

		ptmx, err := pty.Start(cmd)
		if err != nil {
			return err
		}
		defer ptmx.Close()

		io.Copy(stdoutWriter, ptmx)
		fmt.Printf("captured: \n%s\n", stdout.String())

		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
