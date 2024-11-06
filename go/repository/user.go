package repository

import "database/sql"

type AuthRepository interface {
	FindByUsername(username string)
}

type authRepositoryImpl struct {
	db *sql.DB
}
