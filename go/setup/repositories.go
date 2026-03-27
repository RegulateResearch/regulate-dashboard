package setup

import (
	"database/sql"
	"frascati/comp/background"
	"frascati/comp/dbhandler"
	"frascati/comp/txhandler"
	"frascati/repository"
	repo_db "frascati/repository/db"
)

type repositories struct {
	auth         repository.AuthRepository
	course       repository.CourseRepository
	courseMember repository.CourseMemberRepository
	user         repository.UserRepository
	record       repository.RecordRepository
	transactor   txhandler.Transactor
}

func setupRepositories(db *sql.DB, processor background.Processor) repositories {
	executor := dbhandler.NewDbExecutor(db, processor)
	authRepoDb := repo_db.NewAuthRepositoryDb(executor)
	courseRepoDb := repo_db.NewCourseDbRepository(executor)
	courseMemberRepoDb := repo_db.NewCourseMemberDbRepository(executor)
	userRepoDb := repo_db.NewUserRepository(executor)
	recordRepoDb := repo_db.NewRecordRepository(executor)

	return repositories{
		auth:         repository.NewAuthRepository(authRepoDb),
		course:       repository.NewCourseRepository(courseRepoDb),
		courseMember: repository.NewCourseMemberRepository(courseMemberRepoDb),
		user:         repository.NewUserRepository(userRepoDb),
		record:       repository.NewRecordRepository(recordRepoDb),
		transactor:   executor,
	}
}
