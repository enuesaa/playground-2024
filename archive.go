package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ArchiveCurrentDir() (*bytes.Buffer, error) {
	buf := bytes.NewBuffer([]byte{})

	gz := gzip.NewWriter(buf)
	defer gz.Close()

	t := tar.NewWriter(gz)
	defer t.Close()

	err := filepath.Walk(".", func(path string, finfo os.FileInfo, err error) error {
		if err != nil || path == "." {
			return err
		}

		// metadata
		filename := finfo.Name()
		isDir := finfo.IsDir()

		// add file
		header, err := tar.FileInfoHeader(finfo, filename)
		if err != nil {
			return err
		}
		header.Name = path
		if err := t.WriteHeader(header); err != nil {
			return err
		}
		if isDir {
			return nil
		}

		// write file content
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		
		if _, err = io.Copy(t, f); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return nil, err
	}

	return buf, nil
}
