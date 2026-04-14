package entity

import (
	"frascati/constants"
)

type User struct {
	Base
	Email       string
	Password    string
	Username    string
	DisplayName string
	HasSsoLogin bool
	CivitasID   string
	Role        constants.Role
}

func NewUser() User {
	return User{}
}
