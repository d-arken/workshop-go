package main

import (
	"github.com/d-arken/workshop-go/5_orm/db"
	"github.com/d-arken/workshop-go/5_orm/router"
	"github.com/d-arken/workshop-go/5_orm/user"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	gormCfg := &gorm.Config{}
	dbConn, err := db.NewPostgresSQL(gormCfg).Open()
	if err != nil {
		panic(err.Error())
	}

	err = db.Migrate(dbConn)
	if err != nil {
		panic(err.Error())
	}

	userh := user.NewHandler(user.NewService(user.NewRepository(dbConn)))
	r := router.Setup(userh)

	r.Run(":8080")
}
