package repository

import (
	"io"
	"os/exec"
)

type CmdRepositoryInterface interface {
	Exec(writer io.Writer, workdir string, command string, args []string) (*exec.Cmd, error)
	Kill(cmd *exec.Cmd) error
}
type CmdRepository struct{}

func (repo *CmdRepository) Exec(writer io.Writer, workdir string, command string, args []string) (*exec.Cmd, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = workdir

	cmd.Stdout = writer
	cmd.Stderr = writer

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (repo *CmdRepository) Kill(cmd *exec.Cmd) error {
	return cmd.Process.Kill() 
}
