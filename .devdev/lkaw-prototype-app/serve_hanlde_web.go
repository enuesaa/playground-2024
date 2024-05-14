package main

import (
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/enuesaa/lkaw/repository"
)

func HanldeWeb(repos repository.Repos, cli Cli) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
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
	}
	return handler
}
