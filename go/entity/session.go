package entity

import (
	"frascati/constants"
	"frascati/typing"
)

type SessionData struct {
	ID   typing.ID      `json:"id"`
	Role constants.Role `json:"role"`
}
