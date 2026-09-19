package dto

import (
	"frascati/typing"
	"time"
)

type Base struct {
	ID        typing.ID `json:"id"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
}
