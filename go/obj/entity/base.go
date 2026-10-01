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

func newBase() Base {
	return Base{}
}

func baseWithID(id typing.ID) Base {
	return Base{ID: id}
}
