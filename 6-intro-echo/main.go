package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id    int    `json: "id" from:"id"`
	Age   int    `json: "age" from:"age"`
	Name  string `json: "name" from:"name"`
	Email string `json: "email" from:"email"`
}

func main() {
	e := echo.New()
	e.GET("/user/:id/:age", UserController)
	e.POST("/user", RegisterController)
	e.Start(":8000")
}

func UserController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	age, _ := strconv.Atoi(e.Param("age"))

	search := e.QueryParam("search")
	short := e.QueryParam("short")

	user := User{Id: id, Age: age, Name: "Alta", Email: "academy@alterra.id"}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"user":   user,
		"search": search,
		"short":  short,
	})
}

func RegisterController(c echo.Context) error {
	// email := c.FormValue("email")
	// name := c.FormValue("name")

	user := User{}
	c.Bind(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
