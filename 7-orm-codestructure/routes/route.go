package routes

import (
	"project/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	return e
}
