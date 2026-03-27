package repository

import (
	"frascati/exception"
	"frascati/obj/entity"
	repo_db "frascati/repository/db"
	"frascati/typing"
)

type CourseMemberRepository interface {
	FindByCourse(ctx typing.Context, course entity.Course) ([]entity.CourseMember, exception.Exception)
	AddMultiple(ctx typing.Context, newMembers []entity.CourseMember) ([]entity.CourseMember, exception.Exception)
}

type courseMemberRepositoryImpl struct {
	repoDb repo_db.CourseMemberRepository
}

func NewCourseMemberRepository(repoDb repo_db.CourseMemberRepository) CourseMemberRepository {
	return courseMemberRepositoryImpl{
		repoDb: repoDb,
	}
}

func (r courseMemberRepositoryImpl) FindByCourse(ctx typing.Context, course entity.Course) ([]entity.CourseMember, exception.Exception) {
	res, err := r.repoDb.FindByCourse(ctx, course)
	return res, err
}

func (r courseMemberRepositoryImpl) AddMultiple(ctx typing.Context, newMembers []entity.CourseMember) ([]entity.CourseMember, exception.Exception) {
	res, err := r.repoDb.AddMultiple(ctx, newMembers)
	return res, err
}
