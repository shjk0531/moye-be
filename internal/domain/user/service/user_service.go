package service

import (
	"github.com/shjk0531/moye-be/internal/domain/user/model"
	"github.com/shjk0531/moye-be/internal/domain/user/repository"
)

type Service interface {
	RegisterUser(user *model.User) error
	GetUser(id uint) (*model.User, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) RegisterUser(user *model.User) error {
	return s.repo.Create(user)
}

func (s *service) GetUser(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}
