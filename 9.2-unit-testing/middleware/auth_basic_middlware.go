package middleware

import (
	"fmt"
	"project/config"
	"project/models"

	"github.com/labstack/echo"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var user models.Users
	fmt.Println(email, password)
	if err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
