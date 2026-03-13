package setup

import (
	"database/sql"
	"frascati/comp/dbhandler"
	"frascati/comp/txhandler"
	"frascati/repository"
	repo_db "frascati/repository/db"
)

type repositories struct {
	auth       repository.AuthRepository
	user       repository.UserRepository
	transactor txhandler.Transactor
}

func setupRepositories(db *sql.DB) repositories {
	executor := dbhandler.NewDbExecutor(db)
	authRepoDb := repo_db.NewAuthRepositoryDb(executor)
	userRepoDb := repo_db.NewUserRepository(executor)

	return repositories{
		auth:       repository.NewAuthRepository(authRepoDb),
		user:       repository.NewUserRepository(userRepoDb),
		transactor: executor,
	}
}
