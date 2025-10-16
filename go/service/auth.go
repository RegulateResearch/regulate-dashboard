package service

import (
	"errors"
	"frascati/comp/auth"
	auth_exception "frascati/comp/auth/exception"
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository"

	"frascati/typing"
)

type AuthService interface {
	Register(typing.Context, entity.User) (entity.User, exception.Exception)
	Login(typing.Context, entity.User) (string, exception.Exception)
}

type authServiceImpl struct {
	repo          repository.AuthRepository
	bcryptService auth.BcryptService
	jwtService    auth.JwtService
}

func NewAuthService(userRepo repository.AuthRepository, bcryptService auth.BcryptService, jwtService auth.JwtService) AuthService {
	return authServiceImpl{
		repo:          userRepo,
		bcryptService: bcryptService,
		jwtService:    jwtService,
	}
}

func (s authServiceImpl) Register(ctx typing.Context, userWrite entity.User) (entity.User, exception.Exception) {
	emailExist, err := s.repo.IsExistByEmail(ctx, userWrite.Email)
	if err != nil {
		return entity.User{}, err
	}

	if emailExist {
		return entity.User{}, auth_exception.GenerateErrUserAlreadyExist()
	}

	plainPassword := userWrite.Password

	hashedPassword, err := s.bcryptService.HashPassword(plainPassword)
	if err != nil {
		return entity.User{}, err
	}

	newUserData := entity.User{
		Email:    userWrite.Email,
		Username: userWrite.Username,
		Password: string(hashedPassword),
	}

	user, err := s.repo.Add(ctx, newUserData)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s authServiceImpl) Login(ctx typing.Context, userWrite entity.User) (string, exception.Exception) {
	user, err := s.repo.FindByEmail(ctx, userWrite.Email)
	if err != nil {
		if err.Cause() == exception.CAUSE_NOT_FOUND {
			return "", auth_exception.GenerateErrLoginFail(err)
		}

		return "", auth_exception.GenerateErrAuthFailComposite(err)
	}

	loginSuccess := s.bcryptService.ComparePassword(user.Password, userWrite.Password)
	if !loginSuccess {
		return "", auth_exception.GenerateErrLoginFail(errors.New("wrong password"))
	}

	sessionData := entity.Session{
		ID:   user.ID,
		Role: user.Role,
	}

	tokenStr, err := s.jwtService.GenerateToken(sessionData)
	if err != nil {
		return "", auth_exception.GenerateErrAuthFailComposite(err)
	}

	return tokenStr, nil
}
