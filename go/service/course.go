package service

import (
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type CourseService interface {
	FindAll(ctx typing.Context) ([]entity.Course, exception.Exception)
	Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception)
}

type courseServiceImpl struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return courseServiceImpl{
		repo: repo,
	}
}

func (s courseServiceImpl) FindAll(ctx typing.Context) ([]entity.Course, exception.Exception) {
	res, err := s.repo.FindAll(ctx)
	return res, err
}

func (s courseServiceImpl) Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception) {
	res, err := s.repo.Add(ctx, course)
	return res, err
}
