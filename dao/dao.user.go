package dao

import (
	"github.com/renaldyhidayatt/crud_blog/dto"
)

type DaoUser interface {
	GetAll() ([]dto.Users, error)
	GetID(id int) (dto.Users, error)
	Insert(input dto.Users) (dto.Users, error)
	Update(input dto.Users) (dto.Users, error)
	Delete(input int64) error
}
