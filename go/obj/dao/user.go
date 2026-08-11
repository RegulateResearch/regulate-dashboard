package dao

import (
	"database/sql"
	"frascati/constants"
)

// should not be used outside repo_db
type UserDb struct {
	BaseDb
	Email        sql.NullString
	Password     string
	Username     string
	DisplayName  string
	HasSsoLogin  bool
	CivitasID    sql.NullString
	Role         constants.Role
	AcademicRole constants.AcademicRole
}

func NewUserDb() UserDb {
	return UserDb{
		BaseDb: newBaseDb(),
	}
}
