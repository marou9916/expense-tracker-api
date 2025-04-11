package models

type LoginInput struct {
	Email    string `json:"email_user" binding:"required, email"`
	Password string `json:"password_user" binding:"required"`
}

type RegisterInput struct {
	Name     string `json:"name_user" binding:"required, email"`
	Email    string `json:"email_user" binding:"required"`
	Password string `json:"password_user" binding:"required"`
}
