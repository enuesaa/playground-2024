package repository

import (
	"io"
	"os"
)

type FsRepositoryInterface interface {
	IsExist(path string) bool
	IsDir(path string) (bool, error)
	WorkDir() (string, error)
	Read(path string) ([]byte, error)
	Create(path string, body io.Reader) error
	CreateDir(path string) error
}
type FsRepository struct{}

func (repo *FsRepository) IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *FsRepository) IsDir(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return f.IsDir(), nil
}

func (repo *FsRepository) WorkDir() (string, error) {
	return os.Getwd()
}

func (repo *FsRepository) Read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	return io.ReadAll(f)
}

func (repo *FsRepository) Create(path string, body io.Reader) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := io.Copy(file, body); err != nil {
		return err
	}
	return nil
}

func (repo *FsRepository) CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
