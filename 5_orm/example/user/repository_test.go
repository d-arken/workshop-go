package user_test

import (
	"github.com/d-arken/workshop-go/5_orm/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

import . "github.com/d-arken/workshop-go/5_orm/user"

func TestRepository_Create(t *testing.T) {
	dbConn, _ := db.NewSQLiteInMemory().Open()
	_ = db.Migrate(dbConn)
	repo := NewRepository(dbConn)

	user := Model{Name: "Testenilson Silva", Age: 65}
	err := repo.Create(&user)

	assert.Nil(t, err)
}
