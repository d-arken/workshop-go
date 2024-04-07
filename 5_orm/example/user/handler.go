package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Svc ServiceInterface
}

func NewHandler(svc ServiceInterface) *Handler {
	return &Handler{Svc: svc}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user CreateUserRequest
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.Svc.Create(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
