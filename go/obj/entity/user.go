package entity

import (
	"frascati/constants"
	"frascati/typing"
)

type User struct {
	Base
	Email        string
	Password     string
	Username     string
	DisplayName  string
	HasSsoLogin  bool
	CivitasID    string
	Role         constants.Role
	AcademicRole constants.AcademicRole
}

func NewUser() User {
	return User{}
}

func UserWithID(id typing.ID) User {
	return User{Base: baseWithID(id)}
}
