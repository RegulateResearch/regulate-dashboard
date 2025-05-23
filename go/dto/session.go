package dto

import "frascati/typing"

type SessionData struct {
	ID   typing.ID `json:"id"`
	Role string    `json:"role"`
}
