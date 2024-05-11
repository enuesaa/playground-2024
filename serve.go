package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/enuesaa/lkaw/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type FileResponse struct {
	Files []File `json:"files"`
}
type File struct {
	Filename string `json:"filename"`
	Exists bool `json:"exists"`
	Content string `json:"content"`
}

func Serve(repos repository.Repos, cli CLI) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001"},
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/api/files", func(w http.ResponseWriter, r *http.Request) {
		filenames := cli.Filenames
		files := make([]File, 0)
		for _, filename := range filenames {
			content, err := repos.Fs.Read(filename)
			if err != nil {
				files = append(files, File{
					Filename: filename,
					Exists: false,
					Content: "",
				})
				continue
			}
			files = append(files, File{
				Filename: filename,
				Exists: true,
				Content: string(content),
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

	addr := fmt.Sprintf(":%d", cli.Port)
	return http.ListenAndServe(addr, r)
}
