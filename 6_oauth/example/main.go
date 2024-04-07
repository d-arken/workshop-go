package main

import (
	"encoding/gob"
	"github.com/coreos/go-oidc"
	"github.com/d-arken/workshop-go/6_oauth/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Use(auth.JWTProtected()).GET("/authenticated-ping", func(c *gin.Context) {
		claims := c.MustGet(auth.OIDCClaimsContext).(auth.OIDCClaims)
		c.JSON(200, gin.H{"claims": claims})
	})
	return r
}

func main() {
	err := godotenv.Load()
	gob.Register(oidc.IDToken{})
	if err != nil {
		panic(err.Error())
	}

	r := setupRouter()
	r.Run(":8080")
}
