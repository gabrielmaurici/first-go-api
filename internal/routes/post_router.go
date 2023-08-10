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
		err := Create(rp, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusCreated)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		post, err := Get(rp, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.Header().Add("Content-Type", "json")
		w.WriteHeader(http.StatusOK)
		w.Write(post)
	})

	r.Put("/", Update)

	return r
}

func Create(rp *RouterPost, r *http.Request) error {
	var dto usecase.PostInputDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return err
	}

	err = rp.PostUsecase.Create(dto)
	if err != nil {
		return err
	}

	return nil
}

func Get(rp *RouterPost, r *http.Request) ([]byte, error) {
	id := chi.URLParam(r, "id")

	post, err := rp.PostUsecase.Get(id)
	if err != nil {
		return nil, err
	}

	postJson, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	return postJson, nil
}

func Update(w http.ResponseWriter, r *http.Request) {
	//TODO
}
