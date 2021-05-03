package main

import (
	"project/config"
	"project/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8000")
}
