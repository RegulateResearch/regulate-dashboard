package entity

import (
	"frascati/typing"
	"time"
)

type Base struct {
	ID        typing.ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
