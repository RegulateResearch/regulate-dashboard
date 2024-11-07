package dto

import "frascati/constant"

type User struct {
	Base
	Username string        `json:"username"`
	Role     constant.Role `json:"role"`
}

type UserTokenReturn struct {
	ID   int64         `json:"id"`
	Role constant.Role `json:"role"`
}

type UserWrite struct {
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password" binding:"required"`
}
