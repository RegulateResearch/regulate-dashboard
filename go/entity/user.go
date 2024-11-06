package entity

import "frascati/constant"

type User struct {
	Base
	Username string
	Password string
	Role     constant.Role
}

type UserWrite struct {
	Username string
	Password string
}
