package repository

import (
	"errors"
	"frascati/exception"
	"frascati/obj/entity"
	repo_db "frascati/repository/db"
	"frascati/typing"
)

type CourseRepository interface {
	FindAll(ctx typing.Context) ([]entity.Course, exception.Exception)
	FindById(ctx typing.Context, id typing.ID) (entity.Course, exception.Exception)
	FindAllByEnrollingUserId(ctx typing.Context, user entity.User) ([]entity.Course, exception.Exception)
	Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception)
	UpdateById(ctx typing.Context, id typing.ID, updateDate entity.Course) (bool, exception.Exception)
	DeleteById(ctx typing.Context, id typing.ID) (bool, exception.Exception)
	IsExistById(ctx typing.Context, id typing.ID) (bool, exception.Exception)
	CheckExistById(ctx typing.Context, id typing.ID) exception.Exception
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

func (r courseRepositoryImpl) FindAllByEnrollingUserId(ctx typing.Context, user entity.User) ([]entity.Course, exception.Exception) {
	res, err := r.dbRepo.FindAllByEnrollingUserId(ctx, user)
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

func (r courseRepositoryImpl) IsExistById(ctx typing.Context, id typing.ID) (bool, exception.Exception) {
	res, err := r.dbRepo.IsExistById(ctx, id)
	return res, err
}

func (r courseRepositoryImpl) CheckExistById(ctx typing.Context, id typing.ID) exception.Exception {
	isExist, err := r.IsExistById(ctx, id)
	if err != nil {
		return err
	}

	if !isExist {
		return exception.NewBaseException(
			exception.CAUSE_NOT_FOUND,
			"course_member/repository",
			"record not found",
			errors.New("course not found"),
		)
	}

	return nil
}
