package repository

import (
	"frascati/exception"
	"frascati/obj/entity"
	repo_db "frascati/repository/db"
	"frascati/typing"
)

type CourseRepository interface {
	FindAll(ctx typing.Context) ([]entity.Course, exception.Exception)
	FindById(ctx typing.Context, id typing.ID) (entity.Course, exception.Exception)
	Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception)
	UpdateById(ctx typing.Context, id typing.ID, updateDate entity.Course) (bool, exception.Exception)
	DeleteById(ctx typing.Context, id typing.ID) (bool, exception.Exception)
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

func (r courseRepositoryImpl) FindById(ctx typing.Context, id typing.ID) (entity.Course, exception.Exception) {
	res, err := r.dbRepo.FindById(ctx, id)
	return res, err
}

func (r courseRepositoryImpl) UpdateById(ctx typing.Context, id typing.ID, updateDate entity.Course) (bool, exception.Exception) {
	res, err := r.dbRepo.UpdateById(ctx, id, updateDate)
	return res, err
}

func (r courseRepositoryImpl) DeleteById(ctx typing.Context, id typing.ID) (bool, exception.Exception) {
	res, err := r.dbRepo.DeleteById(ctx, id)
	return res, err
}

func (r courseRepositoryImpl) Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception) {
	res, err := r.dbRepo.Add(ctx, course)
	return res, err
}
