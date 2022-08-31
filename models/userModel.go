package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	LastLogin time.Time `json:"last_login"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Username  string    `gorm:"size:200;unique" json:"username"`
	Password  string
	BirthDate string `json:"birth_date"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	Role      string `gorm:"size:200"`
}

type UserSignIn struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
}

type PrepareUser struct {
	Username string
	Email    string
	Role     string
}
