package gateway

import "gabrielmaurici/first-go-api/internal/entity"

type PostGateway interface {
	Save(post *entity.Post) error
	Update(post *entity.Post) error
	Get(id *string) (*entity.Post, error)
	GetAll(offset *string, limit *string) ([]*entity.Post, error)
}
