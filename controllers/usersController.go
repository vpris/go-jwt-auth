package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vpris/test-jwt/initializers"
	"github.com/vpris/test-jwt/models"
	"github.com/vpris/test-jwt/utils"
)

func Signup(c *gin.Context) {
	//Get the email/pass off req body
	var body struct {
		Email    string
		Password string
		Username string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	//Hash the password
	hash := utils.HashPassword(body.Password)

	//Create the user
	user := models.User{Email: body.Email, Password: string(hash), Username: body.Username, Role: "user"}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}
	//Respond
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}

func Login(c *gin.Context) {
	//Get the email and pass off req body
	var body models.UserSignIn

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
	}

	err := utils.ComparePassword(user.Password, body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	gen, err := generateTokenPair(user.ID, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	//Update last login
	initializers.DB.Model(&user).Update("LastLogin", time.Now())

	//Send in back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("access_token", gen["access_token"], 3600*24*30, "", "", false, true)
	c.SetCookie("refresh_token", gen["refresh_token"], 3600*24, "", "", false, true)
	c.SetCookie("logged_in", "true", 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"access_token": gen["access_token"],
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "", "", false, true)
	c.SetCookie("refresh_token", "", -1, "", "", false, true)
	c.SetCookie("logged_in", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
