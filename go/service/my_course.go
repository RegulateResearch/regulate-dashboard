package service

import (
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type myCourseService interface {
	MyCourses(ctx typing.Context, userID typing.ID) ([]entity.Course, exception.Exception)
}

type myCourseServiceImpl struct {
	myBaseService
	courseRepo repository.CourseRepository
}

func newMyCourseService(base myBaseService, courseRepo repository.CourseRepository) myCourseService {
	return myCourseServiceImpl{
		myBaseService: base,
		courseRepo:    courseRepo,
	}
}

func (s myCourseServiceImpl) MyCourses(ctx typing.Context, userID typing.ID) ([]entity.Course, exception.Exception) {
	checkErr := s.checkUserExistByID(ctx, userID)
	if checkErr != nil {
		return nil, checkErr
	}

	user := entity.User{Base: entity.Base{ID: userID}}

	res, err := s.courseRepo.FindAllByEnrollingUserId(ctx, user)
	return res, err
}
