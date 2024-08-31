package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/erikgeiser/promptkit/textinput"
)

func init() {
	log.SetFlags(0)
}

type LogRepositoryInterface interface {
	Info(format string, v ...any)
	Fatal(err error)
	Ask(message string, defaultValue string) (string, error)
	Confirm(message string) (bool, error)
}
type LogRepository struct{}

func (repo *LogRepository) prefix() string {
	return time.Now().Local().Format("15:04:05")
}

func (repo *LogRepository) Info(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	log.Printf("%s  %s\n", repo.prefix(), message)
}

func (repo *LogRepository) Fatal(err error) {
	log.Fatalf("%s  Error: %s\n", repo.prefix(), err.Error())
}

func (repo *LogRepository) Ask(message string, defaultValue string) (string, error) {
	input := textinput.New(message)
	input.InitialValue = defaultValue

	return input.RunPrompt()
}

func (repo *LogRepository) Confirm(message string) (bool, error) {
	message = fmt.Sprintf("%s  %s (y/n)", repo.prefix(), message)
	input := textinput.New(message)
	input.InitialValue = ""

	answer, err := input.RunPrompt()
	if err != nil {
		return false, err
	}

	return answer == "y", nil
}
