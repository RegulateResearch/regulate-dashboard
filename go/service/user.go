package service

import (
	"context"
	"frascati/entity"
	"frascati/exception"
	"frascati/repository"
)

type UserService interface {
	FindAll(ctx context.Context) ([]entity.User, exception.Exception)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userServiceImpl{
		userRepo: userRepo,
	}
}

func (s userServiceImpl) FindAll(ctx context.Context) ([]entity.User, exception.Exception) {
	res, err := s.userRepo.FindAll(ctx)
	return res, err
}
