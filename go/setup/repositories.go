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
	course     repository.CourseRepository
	user       repository.UserRepository
	transactor txhandler.Transactor
}

func setupRepositories(db *sql.DB) repositories {
	executor := dbhandler.NewDbExecutor(db)
	authRepoDb := repo_db.NewAuthRepositoryDb(executor)
	courseRepoDb := repo_db.NewCourseDbRepository(executor)
	userRepoDb := repo_db.NewUserRepository(executor)

	return repositories{
		auth:       repository.NewAuthRepository(authRepoDb),
		course:     repository.NewCourseRepository(courseRepoDb),
		user:       repository.NewUserRepository(userRepoDb),
		transactor: executor,
	}
}
