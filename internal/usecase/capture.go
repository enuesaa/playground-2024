package usecase

import (
	"bytes"
	"image/png"

	"github.com/enuesaa/codetrailer/internal/repository"
)

func Capture(repos repository.Repos, path string) error {
	img, err := repos.Fs.Capture()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}
	return repos.Fs.Create(path, &buf)
}
