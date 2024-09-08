package repository

import (
	"io"
	"os/exec"
)

type CmdRepositoryInterface interface {
	Exec(command string, writer io.Writer) error
}
type CmdRepository struct{}

func (repo *CmdRepository) Exec(command string, writer io.Writer) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = writer
	cmd.Stderr = writer

	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}
