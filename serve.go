package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Serve(port int) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001"},
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/api/files", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"a":"b"}`))
	})

	addr := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(addr, r)
}
