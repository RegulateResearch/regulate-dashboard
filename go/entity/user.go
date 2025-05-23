package entity

import "frascati/constants"

type User struct {
	Base
	Username string
	Password string
	Role     constants.Role
}

type UserWrite struct {
	Username string
	Password string
}
