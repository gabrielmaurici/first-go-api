package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouterPost() http.Handler {
	r := chi.NewRouter()

	r.Post("/", create)
	r.Get("/", get)
	r.Put("/", update)

	return r
}

func create(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func get(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func update(w http.ResponseWriter, r *http.Request) {
	//TODO
}
