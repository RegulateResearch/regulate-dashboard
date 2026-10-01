package service

import (
	"frascati/comp/txhandler"
	"frascati/repository"
)

type MyService interface {
	myCourseService
	myProfileService
	myTaskService
}

type myServiceImpl struct {
	myCourseService
	myProfileService
	myTaskService
}

func NewMyService(
	userRepo repository.UserRepository,
	courseRepo repository.CourseRepository,
	userTaskRepo repository.UserTaskRepository,
	taskRecordRepo repository.TaskRecordRepository,
	transactor txhandler.Transactor,
) MyService {
	base := newMyBaseService(userRepo)
	return myServiceImpl{
		myCourseService:  newMyCourseService(base, courseRepo),
		myProfileService: newMyProfileService(base, userRepo),
		myTaskService:    newMyTaskService(base, userTaskRepo, taskRecordRepo, transactor),
	}
}
