package main

import (
	"encoding/json"
	"net/http"

	"github.com/enuesaa/lkaw/repository"
)

type FileResponse struct {
	Files []File `json:"files"`
}
type File struct {
	Filename string `json:"filename"`
	Exists   bool   `json:"exists"`
	Content  string `json:"content"`
}

func HanldeFiles(repos repository.Repos, cli Cli) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		files := make([]File, 0)
		for _, filename := range cli.Filenames {
			isdir, err := repos.Fs.IsDir(filename)
			if err != nil || isdir {
				continue
			}
			content, err := repos.Fs.Read(filename)
			if err != nil {
				files = append(files, File{
					Filename: filename,
					Exists:   false,
					Content:  "",
				})
				continue
			}
			files = append(files, File{
				Filename: filename,
				Exists:   true,
				Content:  string(content),
			})
		}
		response := FileResponse{
			Files: files,
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(responseBytes)
	}
	return handler
}
