package repository

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

type FsMockRepository struct {
	Files []string
}

func (repo *FsMockRepository) Ext(path string) string {
	return filepath.Ext(path)
}

func (repo *FsMockRepository) IsExist(path string) bool {
	return slices.Contains(repo.Files, path)
}

func (repo *FsMockRepository) IsDir(path string) (bool, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/workdir/" + path
	}
	result := new(bool)
	for _, filepath := range repo.Files {
		if strings.HasPrefix(filepath, path) {
			if filepath == path {
				// this is file. not dir.
				if *result != true {
					*result = false
				}
				continue
			}
			*result = true
			continue
		}
	}
	if result != nil {
		return *result, nil
	}

	return false, fmt.Errorf("file or dir does not exist.")
}

func (repo *FsMockRepository) WorkDir() (string, error) {
	return "/workdir", nil
}

func (repo *FsMockRepository) Read(path string) ([]byte, error) {
	return make([]byte, 0), nil
}
