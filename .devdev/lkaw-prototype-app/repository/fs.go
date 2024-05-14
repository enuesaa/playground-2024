package repository

import (
	"io"
	"os"
	"path/filepath"
)

type FsRepositoryInterface interface {
	IsDir(path string) (bool, error)
	Read(path string) ([]byte, error)
	ListFiles(path string) ([]string, error)
}
type FsRepository struct{}

func (repo *FsRepository) IsDir(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return f.IsDir(), nil
}

func (repo *FsRepository) Read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	return io.ReadAll(f)
}

func (repo *FsRepository) ListFiles(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	filenames := make([]string, 0)
	for _, entry := range entries {
		if entry.Name() == ".git" {
			continue
		}
		path := filepath.Join(path, entry.Name())
		filenames = append(filenames, path)
	}
	return filenames, nil
}
