package services

import (
	"github.com/renaldyhidayatt/crud_blog/dao"
	"github.com/renaldyhidayatt/crud_blog/dto"
)

type userService struct {
	user dao.DaoUser
}

func NewUserService(user dao.DaoUser) *userService {
	return &userService{user: user}
}

func (s *userService) GetAll() ([]dto.Users, error) {
	res, err := s.user.GetAll()

	return res, err
}

func (s *userService) GetID(id int) (dto.Users, error) {
	res, err := s.user.GetID(id)

	return res, err
}

func (s *userService) Insert(usr dto.Users) (dto.Users, error) {
	var user dto.Users

	user.Name = usr.Name
	user.Hobby = usr.Hobby
	user.CreatedAt = usr.CreatedAt
	user.UpdatedAt = usr.UpdatedAt

	res, err := s.user.Insert(user)

	return res, err

}

func (s *userService) Update(usr dto.Users) (dto.Users, error) {
	var user dto.Users

	user.ID = usr.ID
	user.Name = usr.Name
	user.Hobby = usr.Hobby
	user.CreatedAt = usr.CreatedAt
	user.UpdatedAt = usr.UpdatedAt
	res, err := s.user.Update(user)

	return res, err
}

func (s *userService) Delete(id int64) error {
	err := s.user.Delete(id)

	return err
}
