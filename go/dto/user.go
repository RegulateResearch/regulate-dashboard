package dto

import "frascati/constants"

type User struct {
	Base
	Username string         `json:"username"`
	Role     constants.Role `json:"role"`
}

type UserTokenReturn struct {
	ID   int64          `json:"id"`
	Role constants.Role `json:"role"`
}

type UserWrite struct {
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password" binding:"required"`
}
