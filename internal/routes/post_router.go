package routes

import (
	"encoding/json"
	"gabrielmaurici/first-go-api/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RouterPost struct {
	PostUsecase usecase.PostUseCase
}

func NewRouterPost(postUsecase *usecase.PostUseCase) *RouterPost {
	return &RouterPost{
		PostUsecase: *postUsecase,
	}
}

func (rp *RouterPost) AddHandlerPost() http.Handler {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var dto usecase.PostInputDto

		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		rp.PostUsecase.Create(dto)
	})

	r.Get("/", Get)
	r.Put("/", Update)

	return r
}

func Get(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func Update(w http.ResponseWriter, r *http.Request) {
	//TODO
}
