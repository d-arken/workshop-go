package main

import (
	"fmt"
	"github.com/d-arken/workshop-go/7_crawler/pokemon"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	_ = godotenv.Load()
	r := gin.Default()
	r.Use(CORSMiddleware())

	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_TZ"))), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	_ = db.AutoMigrate(pokemon.Model{})

	handler := pokemon.Handler{
		Repo: &pokemon.Repository{
			DB: db,
		},
	}

	r.GET("/list", handler.PopulatePokemonsList)
	r.GET("/populate", handler.PopulatePokemons)

	r.Run(":8080")

}
