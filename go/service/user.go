package service

import (
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type UserService interface {
	FindAll(ctx typing.Context) ([]entity.User, exception.Exception)
	FindById(ctx typing.Context, id typing.ID) (entity.User, exception.Exception)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userServiceImpl{
		userRepo: userRepo,
	}
}

func (s userServiceImpl) FindAll(ctx typing.Context) ([]entity.User, exception.Exception) {
	res, err := s.userRepo.FindAll(ctx)
	return res, err
}

func (s userServiceImpl) FindById(ctx typing.Context, id typing.ID) (entity.User, exception.Exception) {
	res, err := s.userRepo.FindById(ctx, id)
	return res, err
}
