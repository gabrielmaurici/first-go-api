package gateway

import "gabrielmaurici/first-go-api/internal/entity"

type PostGateway interface {
	Get(id *string) (*entity.Post, error)
	Save(post *entity.Post) error
	Update(post *entity.Post) error
}
