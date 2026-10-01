package dao

import (
	"database/sql"
	"frascati/typing"
	"time"
)

// only for scanning
// should not be used outside repo_db
type BaseDb struct {
	ID        typing.ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func newBaseDb() BaseDb {
	return BaseDb{}
}
