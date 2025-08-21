package dto

import "frascati/typing"

type Session struct {
	ID   typing.ID `json:"id"`
	Role string    `json:"role"`
}
