package router

import (
	"github.com/d-arken/workshop-go/5_orm/user"
	"github.com/gin-gonic/gin"
)

func Setup(userh *user.Handler) *gin.Engine {
	r := gin.Default()
	r.POST("/", userh.CreateUser)
	return r
}
