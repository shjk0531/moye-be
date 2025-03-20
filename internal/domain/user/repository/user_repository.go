package repository

import (
	"github.com/shjk0531/moye-be/internal/domain/user/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.User) error
	FindByID(id uint) (*model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
