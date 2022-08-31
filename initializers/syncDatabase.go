package initializers

import (
	"fmt"
	"os"

	"github.com/vpris/test-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func SyncDatabase() {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASS")), 10)

	if err != nil {
		fmt.Println("Failed to hash password")
		return
	}

	var users = []models.User{{
		Email:     "vpris@example.ru",
		Username:  "vpris",
		Password:  string(passwordHash),
		BirthDate: string("01.01.1991"),
		Height:    177,
		Weight:    85,
		Role:      string("admin"),
	}}

	DB.AutoMigrate(&models.User{})
	DB.Create(&users)
}
