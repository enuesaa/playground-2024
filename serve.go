package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/enuesaa/lkaw/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type FileResponse struct {
	Files []File `json:"files"`
}
type File struct {
	Filename string `json:"filename"`
	Exists   bool   `json:"exists"`
	Content  string `json:"content"`
}

//go:embed dist/*
var dist embed.FS

func Serve(repos repository.Repos, cli Cli) error {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001"},
	}))

	r.Get("/api/files", func(w http.ResponseWriter, r *http.Request) {
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
	})
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		path := fmt.Sprintf("dist%s", r.URL.Path)
		if strings.HasSuffix(path, "/") {
			path += "index.html"
		}

		f, err := dist.ReadFile(path)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		fileExt := filepath.Ext(path)
		mimeType := mime.TypeByExtension(fileExt)
		w.Header().Add("Content-Type", mimeType)
		w.Write(f)
	})

	addr := fmt.Sprintf("localhost:%d", cli.Port)
	fmt.Printf("Serving on %s\n", addr)
	return http.ListenAndServe(addr, r)
}
