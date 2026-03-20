package service

import (
	"errors"
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type CourseService interface {
	FindAll(ctx typing.Context) ([]entity.Course, exception.Exception)
	Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception)
	FindById(ctx typing.Context, id typing.ID) (entity.Course, exception.Exception)
	UpdateById(ctx typing.Context, id typing.ID, updateData entity.Course) exception.Exception
	DeleteById(ctx typing.Context, id typing.ID) exception.Exception
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

func (s courseServiceImpl) FindById(ctx typing.Context, id typing.ID) (entity.Course, exception.Exception) {
	res, err := s.repo.FindById(ctx, id)
	return res, err
}

func (s courseServiceImpl) UpdateById(ctx typing.Context, id typing.ID, updateData entity.Course) exception.Exception {
	success, err := s.repo.UpdateById(ctx, id, updateData)
	if err != nil {
		return err
	}

	if !success {
		return exception.NewBaseException(exception.CAUSE_NOT_FOUND, "course/service", exception.NOT_FOUND, errors.New("update fail"))
	}

	return nil
}

func (s courseServiceImpl) DeleteById(ctx typing.Context, id typing.ID) exception.Exception {
	success, err := s.repo.DeleteById(ctx, id)
	if err != nil {
		return err
	}

	if !success {
		return exception.NewBaseException(exception.CAUSE_NOT_FOUND, "course/service", exception.NOT_FOUND, errors.New("delete fail"))
	}

	return nil
}
