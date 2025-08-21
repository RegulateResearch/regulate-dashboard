package entity

import "frascati/constants"

type User struct {
	Base
	Email    string
	Password string
	Username string
	Role     constants.Role
}
