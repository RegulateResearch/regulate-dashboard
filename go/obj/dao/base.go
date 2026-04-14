package dao

import (
	"frascati/typing"
	"time"
)

// only for scanning
// should not be used outside repo_db
type BaseDb struct {
	ID        typing.ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func newBaseDb() BaseDb {
	return BaseDb{}
}
