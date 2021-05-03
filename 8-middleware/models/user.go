package models

type Users struct {
	ID       int    `json:"id" form:"name"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json: "password" form: "password"`
	Token    string `json: "token" form: "token"`
}
