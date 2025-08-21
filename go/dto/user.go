package dto

import "frascati/constants"

type User struct {
	Base
	Email    string         `json:"email"`
	Username string         `json:"username"`
	Role     constants.Role `json:"role"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}
