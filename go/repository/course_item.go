package repository

import (
	"frascati/exception"
	"frascati/obj/entity"
	repo_db "frascati/repository/db"
	"frascati/typing"
)

type CourseItemRepository interface {
	FindByCourse(ctx typing.Context, course entity.Course) ([]entity.CourseItem, exception.Exception)
	AddBulk(ctx typing.Context, items []entity.CourseItem) ([]entity.CourseItem, exception.Exception)
}

type courseItemRepositoryImpl struct {
	repoDb repo_db.CourseItemRepository
}

func NewCourseItemRepository(repoDb repo_db.CourseItemRepository) CourseItemRepository {
	return courseItemRepositoryImpl{
		repoDb: repoDb,
	}
}

func (r courseItemRepositoryImpl) FindByCourse(ctx typing.Context, course entity.Course) ([]entity.CourseItem, exception.Exception) {
	res, err := r.repoDb.FindByCourse(ctx, course)
	return res, err
}

func (r courseItemRepositoryImpl) AddBulk(ctx typing.Context, items []entity.CourseItem) ([]entity.CourseItem, exception.Exception) {
	res, err := r.repoDb.AddBulk(ctx, items)
	return res, err
}
