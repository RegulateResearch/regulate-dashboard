package dao

import "frascati/constants"

// only for scanning
// should not be used outside repo_db
type UserDb struct {
	BaseDb
	Email    string
	Password string
	Username string
	Role     constants.Role
}

func NewUserDb() UserDb {
	return UserDb{
		BaseDb: newBaseDb(),
	}
}
