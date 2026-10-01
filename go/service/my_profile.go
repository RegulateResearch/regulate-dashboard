package service

import (
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type myProfileService interface {
	MyProfile(ctx typing.Context, userID typing.ID) (entity.User, exception.Exception)
}

type myProfileServiceImpl struct {
	myBaseService
	userRepo repository.UserRepository
}

func newMyProfileService(base myBaseService, userRepo repository.UserRepository) myProfileService {
	return myProfileServiceImpl{
		myBaseService: base,
		userRepo:      userRepo,
	}
}

func (s myProfileServiceImpl) MyProfile(ctx typing.Context, userID typing.ID) (entity.User, exception.Exception) {
	checkErr := s.checkUserExistByID(ctx, userID)
	if checkErr != nil {
		return entity.User{}, checkErr
	}

	user, err := s.userRepo.FindById(ctx, userID)
	return user, err
}
