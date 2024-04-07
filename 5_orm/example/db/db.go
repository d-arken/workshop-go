package db

import (
	"errors"
	"github.com/d-arken/workshop-go/5_orm/user"
	"gorm.io/gorm"
)

type Connector interface {
	Open() (*gorm.DB, error)
}

func Migrate(db *gorm.DB) error {
	var err error
	err = errors.Join(db.AutoMigrate(&user.Model{}))
	return err
}
