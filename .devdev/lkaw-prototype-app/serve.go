package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/enuesaa/lkaw/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/pkg/browser"
)

//go:embed dist/*
var dist embed.FS

func Serve(repos repository.Repos, cli Cli) error {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001"},
	}))

	r.Get("/api/files", HanldeFiles(repos, cli))
	r.Get("/*", HanldeWeb(repos, cli))

	addr := fmt.Sprintf("localhost:%d", cli.Port)
	if err := browser.OpenURL(fmt.Sprintf("http://%s", addr)); err != nil {
		return err
	}

	fmt.Printf("Serving on %s\n", addr)
	return http.ListenAndServe(addr, r)
}
