package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	echomid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUsersController)

	// eBasicAuth := e.Group("")
	// eBasicAuth.Use(echomid.BasicAuth(middleware.BasicAuthDB))
	// eBasicAuth.GET("/users", controllers.GetUserController)

	jwtAuth := e.Group("")
	jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	jwtAuth.GET("/users", controllers.GetUserController)
	return e
}
