package setup

import (
	"database/sql"
	"frascati/repository"
)

type repositories struct {
	auth repository.AuthRepository
	user repository.UserRepository
}

func setupRepositories(db *sql.DB) repositories {
	return repositories{
		auth: repository.NewAuthRepository(db),
		user: repository.NewUserRepository(db),
	}
}
