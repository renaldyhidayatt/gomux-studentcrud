package dao

import (
	"context"

	"github.com/renaldyhidayatt/crud_blog/dto"
)

type DaoUser interface {
	GetAll(ctx context.Context) ([]dto.Users, error)
	GetID(ctx context.Context, id int) (dto.Users, error)
	Insert(ctx context.Context, input *dto.Users) (dto.Users, error)
	Update(ctx context.Context, input dto.Users) (dto.Users, error)
	Delete(ctx context.Context, input int64) error
}
