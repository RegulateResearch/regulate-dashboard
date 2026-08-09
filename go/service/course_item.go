package service

import (
	"frascati/exception"
	"frascati/lambda"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type CourseItemService interface {
	FindByCourseId(ctx typing.Context, courseId typing.ID) ([]entity.CourseItem, exception.Exception)
	AddBulk(ctx typing.Context, courseId typing.ID, items []entity.CourseItem) ([]entity.CourseItem, exception.Exception)
}

type courseItemServiceImpl struct {
	itemRepo   repository.CourseItemRepository
	courseRepo repository.CourseRepository
}

func NewCourseItemService(itemRepo repository.CourseItemRepository, courseRepo repository.CourseRepository) CourseItemService {
	return courseItemServiceImpl{
		itemRepo:   itemRepo,
		courseRepo: courseRepo,
	}
}

func (s courseItemServiceImpl) FindByCourseId(ctx typing.Context, courseId typing.ID) ([]entity.CourseItem, exception.Exception) {
	checkExistErr := s.courseRepo.CheckExistById(ctx, courseId)
	if checkExistErr != nil {
		return nil, checkExistErr
	}

	course := entity.Course{}
	course.ID = courseId

	res, err := s.itemRepo.FindByCourse(ctx, course)
	return res, err
}

func (s courseItemServiceImpl) AddBulk(ctx typing.Context, courseId typing.ID, items []entity.CourseItem) ([]entity.CourseItem, exception.Exception) {
	checkExistErr := s.courseRepo.CheckExistById(ctx, courseId)
	if checkExistErr != nil {
		return nil, checkExistErr
	}

	items = lambda.MapList(items, func(item entity.CourseItem) entity.CourseItem {
		item.Course.ID = courseId
		return item
	})

	res, err := s.itemRepo.AddBulk(ctx, items)
	return res, err
}
