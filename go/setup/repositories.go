package setup

import (
	"database/sql"
	"frascati/repository"
	"frascati/repository/db/queryexec"
)

type repositories struct {
	auth       repository.AuthRepository
	user       repository.UserRepository
	transactor queryexec.Transactor
}

func setupRepositories(db *sql.DB) repositories {
	executor := queryexec.NewDbExecutor(db)

	return repositories{
		auth:       repository.NewAuthRepository(db),
		user:       repository.NewUserRepository(db),
		transactor: executor,
	}
}
