package auth

import (
	"context"
	"frascati/entity"
	"frascati/exception"
	"frascati/repository"
	auth_exception "frascati/service/auth/exception"
)

type AuthService interface {
	Register(context.Context, entity.UserWrite) (entity.User, exception.Exception)
	Login(context.Context, entity.UserWrite) (string, exception.Exception)
}

type authServiceImpl struct {
	repo          repository.AuthRepository
	bcryptService BcryptService
	jwtService    JwtService
}

func NewAuthService(userRepo repository.AuthRepository, bcryptService BcryptService, jwtService JwtService) AuthService {
	return authServiceImpl{
		repo:          userRepo,
		bcryptService: bcryptService,
		jwtService:    jwtService,
	}
}

func (s authServiceImpl) Register(ctx context.Context, userWrite entity.UserWrite) (entity.User, exception.Exception) {
	username := userWrite.Username

	usernameExist, err := s.repo.IsExist(ctx, username)
	if err != nil {
		return entity.User{}, err
	}

	if usernameExist {
		return entity.User{}, auth_exception.GenerateErrUserAlreadyExist()
	}

	plainPassword := userWrite.Password

	hashedPassword, err := s.bcryptService.HashPassword(plainPassword)
	if err != nil {
		return entity.User{}, err
	}

	newUserData := entity.UserWrite{
		Username: username,
		Password: string(hashedPassword),
	}

	user, err := s.repo.Add(ctx, newUserData)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s authServiceImpl) Login(ctx context.Context, userWrite entity.UserWrite) (string, exception.Exception) {
	username := userWrite.Username
	password := userWrite.Password

	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		if err.Cause() == exception.CAUSE_NOT_FOUND {
			return "", auth_exception.GenerateErrLoginFail(err)
		}

		return "", auth_exception.GenerateErrAuthFailComposite(err)
	}

	loginSuccess := s.bcryptService.ComparePassword(user.Password, password)
	if !loginSuccess {
		return "", auth_exception.GenerateErrLoginFail(nil)
	}

	tokenStr, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return "", auth_exception.GenerateErrAuthFailComposite(err)
	}

	return tokenStr, nil
}
