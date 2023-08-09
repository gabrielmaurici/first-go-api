package usecase

import (
	"gabrielmaurici/first-go-api/internal/entity"
	"gabrielmaurici/first-go-api/internal/gateway"
)

type PostInputDto struct {
	Title string
	Body  string
}

type PostUseCase struct {
	PostGateway gateway.PostGateway
}

func NewPostUseCase(postGateway gateway.PostGateway) *PostUseCase {
	return &PostUseCase{
		PostGateway: postGateway,
	}
}

func (uc *PostUseCase) Create(input PostInputDto) error {
	post, err := entity.NewPost(input.Title, input.Body)
	if err != nil {
		return err
	}

	uc.PostGateway.Save(post)

	return nil
}

func (uc *PostUseCase) Get(id string) (*entity.Post, error) {
	post, err := uc.PostGateway.Get(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}
