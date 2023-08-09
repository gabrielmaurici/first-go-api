package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Post struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewPost(title string, body string) (*Post, error) {
	if err := validatePost(title, body); err != nil {
		return nil, errors.New(err.Error())
	}

	return &Post{
		Id:    uuid.New().String(),
		Title: title,
		Body:  body,
	}, nil
}

func (p *Post) UpdatePost(title string, body string) error {
	if err := validatePost(title, body); err != nil {
		return errors.New(err.Error())
	}

	p.Title = title
	p.Body = body

	return nil
}

func validatePost(title string, body string) error {
	if title == "" {
		return errors.New("title e um campo obrigatorio")
	}

	if body == "" {
		return errors.New("body e um campo obrigatorio")
	}

	return nil
}
