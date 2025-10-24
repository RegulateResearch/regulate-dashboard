package setup

import (
	"database/sql"
	"frascati/repository"
	repo_db "frascati/repository/db"
	"frascati/repository/db/queryexec"
)

type repositories struct {
	auth       repository.AuthRepository
	user       repository.UserRepository
	transactor queryexec.Transactor
}

func setupRepositories(db *sql.DB) repositories {
	executor := queryexec.NewDbExecutor(db)
	authRepoDb := repo_db.NewAuthRepositoryDb(executor)
	userRepoDb := repo_db.NewUserRepository(executor)

	return repositories{
		auth:       repository.NewAuthRepository(authRepoDb),
		user:       repository.NewUserRepository(userRepoDb),
		transactor: executor,
	}
}
