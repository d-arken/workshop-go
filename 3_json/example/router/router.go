package router

import (
	"github.com/d-arken/workshop-go/tree/main/3_json/example/user"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.POST("/", user.CreateUser)
	return r
}
