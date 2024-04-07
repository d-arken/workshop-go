package user

import (
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	Create(user *Model) error
}

type Repository struct {
	db *gorm.DB
}

func (r Repository) Create(user *Model) error {
	return r.db.Create(user).Error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
