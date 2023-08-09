package main

import (
	"gabrielmaurici/first-go-api/internal/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Mount("/post", routes.RouterPost())

	http.ListenAndServe(":8080", r)
}
