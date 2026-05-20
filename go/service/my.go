package service

import (
	"errors"
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type MyService interface {
	MyProfile(ctx typing.Context, userID typing.ID) (entity.User, exception.Exception)
	MyCourses(ctx typing.Context, userID typing.ID) ([]entity.Course, exception.Exception)
	checkUserExistByID(ctx typing.Context, userID typing.ID) exception.Exception
}

type myServiceImpl struct {
	userRepo   repository.UserRepository
	courseRepo repository.CourseRepository
}

func NewMyService(userRepo repository.UserRepository, courseRepo repository.CourseRepository) MyService {
	return myServiceImpl{
		userRepo:   userRepo,
		courseRepo: courseRepo,
	}
}

func (s myServiceImpl) MyProfile(ctx typing.Context, userID typing.ID) (entity.User, exception.Exception) {
	checkErr := s.checkUserExistByID(ctx, userID)
	if checkErr != nil {
		return entity.User{}, checkErr
	}

	user, err := s.userRepo.FindById(ctx, userID)
	return user, err
}

func (s myServiceImpl) MyCourses(ctx typing.Context, userID typing.ID) ([]entity.Course, exception.Exception) {
	checkErr := s.checkUserExistByID(ctx, userID)
	if checkErr != nil {
		return nil, checkErr
	}

	user := entity.User{Base: entity.Base{ID: userID}}

	res, err := s.courseRepo.FindAllByEnrollingUserId(ctx, user)
	return res, err
}

func (s myServiceImpl) checkUserExistByID(ctx typing.Context, userID typing.ID) exception.Exception {
	found, err := s.userRepo.IsExistById(ctx, userID)
	if err != nil {
		return err
	}

	if !found {
		baseErr := errors.New("user with such id is not found")
		return exception.NewBaseException(exception.CAUSE_NOT_FOUND, "my/service", "no user with such id", baseErr)
	}

	return nil
}
