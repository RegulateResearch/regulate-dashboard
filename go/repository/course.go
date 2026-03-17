package repository

import (
	"frascati/exception"
	"frascati/obj/entity"
	repo_db "frascati/repository/db"
	"frascati/typing"
)

type CourseRepository interface {
	FindAll(ctx typing.Context) ([]entity.Course, exception.Exception)
	Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception)
}

type courseRepositoryImpl struct {
	dbRepo repo_db.CourseRepository
}

func NewCourseRepository(dbRepo repo_db.CourseRepository) CourseRepository {
	return courseRepositoryImpl{
		dbRepo: dbRepo,
	}
}

func (r courseRepositoryImpl) FindAll(ctx typing.Context) ([]entity.Course, exception.Exception) {
	res, err := r.dbRepo.FindAll(ctx)
	return res, err
}

func (r courseRepositoryImpl) Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception) {
	res, err := r.dbRepo.Add(ctx, course)
	return res, err
}
