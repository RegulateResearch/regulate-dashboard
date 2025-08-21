package entity

import (
	"frascati/constants"
	"frascati/typing"
)

type Session struct {
	ID   typing.ID      `json:"id"`
	Role constants.Role `json:"role"`
}
