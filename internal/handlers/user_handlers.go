package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exist := c.Get("user_id)")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	email, _ := c.Get("email")

	c.JSON(http.StatusOK, gin.H{
		"message": "This is protection endpoint",
		"user_id": userID,
		"email": email,
	})
}