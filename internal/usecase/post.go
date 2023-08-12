package usecase

import (
	"gabrielmaurici/first-go-api/internal/entity"
	"gabrielmaurici/first-go-api/internal/gateway"
)

type PostInputDto struct {
	Title string
	Body  string
}

type PostUpdateDto struct {
	Id    string
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

func (uc *PostUseCase) Create(input *PostInputDto) error {
	post, err := entity.NewPost(input.Title, input.Body)
	if err != nil {
		return err
	}

	err = uc.PostGateway.Save(post)
	if err != nil {
		return err
	}

	return nil
}

func (uc *PostUseCase) Get(id *string) (*entity.Post, error) {
	post, err := uc.PostGateway.Get(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (uc *PostUseCase) Update(input *PostUpdateDto) (*entity.Post, error) {
	post, err := uc.PostGateway.Get(&input.Id)
	if err != nil {
		return nil, err
	}

	err = post.UpdatePost(input.Title, input.Body)
	if err != nil {
		return nil, err
	}

	err = uc.PostGateway.Update(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
