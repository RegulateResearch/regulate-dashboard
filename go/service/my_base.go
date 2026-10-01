package service

import (
	"errors"
	"frascati/exception"
	"frascati/repository"
	"frascati/typing"
)

type myBaseService interface {
	checkUserExistByID(ctx typing.Context, userID typing.ID) exception.Exception
}

type myBaseServiceImpl struct {
	userRepo repository.UserRepository
}

func newMyBaseService(userRepo repository.UserRepository) myBaseService {
	return myBaseServiceImpl{
		userRepo: userRepo,
	}
}

func (s myBaseServiceImpl) checkUserExistByID(ctx typing.Context, userID typing.ID) exception.Exception {
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
