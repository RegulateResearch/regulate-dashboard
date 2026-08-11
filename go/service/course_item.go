package service

import (
	"errors"
	"frascati/exception"
	"frascati/lambda"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type CourseItemService interface {
	FindByCourseId(ctx typing.Context, courseId typing.ID) ([]entity.CourseItem, exception.Exception)
	AddBulk(ctx typing.Context, courseId typing.ID, items []entity.CourseItem) ([]entity.CourseItem, exception.Exception)
	UpdateSingular(ctx typing.Context, courseId typing.ID, itemId typing.ID, item entity.CourseItem) (entity.CourseItem, exception.Exception)
	DeleteSingular(ctx typing.Context, courseId typing.ID, itemId typing.ID) exception.Exception
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

func (s courseItemServiceImpl) UpdateSingular(ctx typing.Context, courseId typing.ID, itemId typing.ID, item entity.CourseItem) (entity.CourseItem, exception.Exception) {
	checkExistErr := s.courseRepo.CheckExistById(ctx, courseId)
	if checkExistErr != nil {
		return entity.CourseItem{}, checkExistErr
	}

	item.ID = itemId
	item.Course.ID = courseId

	res, err := s.itemRepo.UpdateSingular(ctx, item)
	return res, err
}

func (s courseItemServiceImpl) DeleteSingular(ctx typing.Context, courseId typing.ID, itemId typing.ID) exception.Exception {
	checkExistErr := s.courseRepo.CheckExistById(ctx, courseId)
	if checkExistErr != nil {
		return checkExistErr
	}

	item := entity.CourseItem{}
	item.ID = itemId
	item.Course.ID = courseId

	success, err := s.itemRepo.DeleteSingular(ctx, item)
	if err != nil {
		return err
	} else if !success {
		newErr := errors.New("trying to delete nonexistent data")
		return exception.NewBaseException(exception.CAUSE_NOT_FOUND, "course_item", exception.NOT_FOUND, newErr)
	}

	return nil
}
