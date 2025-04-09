package models

type LoginInput struct {
	Email    string	`json:"email_user" binding:"required, email"`
	Password string `json:"password_user" binding:"required"`
}