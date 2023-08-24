package routes

import (
	"encoding/json"
	"errors"
	"fmt"
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

	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		post, err := Update(rp, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.Header().Add("Content-Type", "json")
		w.WriteHeader(http.StatusOK)
		w.Write(post)
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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		post, err := GetAll(rp, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.Header().Add("Content-Type", "json")
		w.WriteHeader(http.StatusOK)
		w.Write(post)
	})

	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		err := Delete(rp, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusNoContent)
	})

	return r
}

func Create(rp *RouterPost, r *http.Request) error {
	var dto usecase.PostInputDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return err
	}

	err = rp.PostUsecase.Create(&dto)
	if err != nil {
		return err
	}

	return nil
}

func Update(rp *RouterPost, r *http.Request) ([]byte, error) {
	var dto usecase.PostUpdateDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return nil, err
	}

	if dto.Id == "" {
		return nil, errors.New("id e um campo obrigatorio")
	}

	post, err := rp.PostUsecase.Update(&dto)
	if err != nil {
		return nil, err
	}

	postJson, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	return postJson, nil
}

func Get(rp *RouterPost, r *http.Request) ([]byte, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, errors.New("id e um campo obrigatorio")
	}

	post, err := rp.PostUsecase.Get(&id)
	if err != nil {
		return nil, err
	}

	postJson, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	return postJson, nil
}

func GetAll(rp *RouterPost, r *http.Request) ([]byte, error) {
	offset := r.URL.Query().Get("offset")
	limit := r.URL.Query().Get("limit")

	fmt.Println("offset: " + offset)
	fmt.Println("limit: " + limit)

	posts, err := rp.PostUsecase.GetAll(&offset, &limit)
	if err != nil {
		return nil, err
	}

	postsJson, err := json.Marshal(posts)
	if err != nil {
		panic(err)
	}

	return postsJson, nil
}

func Delete(rp *RouterPost, r *http.Request) error {
	id := chi.URLParam(r, "id")
	if id == "" {
		return errors.New("id e um campo obrigatorio")
	}

	err := rp.PostUsecase.Delete(&id)
	if err != nil {
		return err
	}

	return nil
}
