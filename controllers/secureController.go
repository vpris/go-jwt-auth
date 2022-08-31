package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vpris/test-jwt/models"
)

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	var prep = models.PrepareUser{
		Email:    user.(models.User).Email,
		Username: user.(models.User).Username,
		Role:     user.(models.User).Role,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello, %v!", prep.Username),
	})
}

func ValidAdmin(c *gin.Context) {
	user, _ := c.Get("user")

	var prep = models.PrepareUser{
		Email:    user.(models.User).Email,
		Username: user.(models.User).Username,
		Role:     user.(models.User).Role,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello, %s! You're admin!", prep.Username),
	})
}
