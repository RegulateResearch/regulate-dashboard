package service

import (
	"fmt"
	"frascati/comp/auth"
	"frascati/comp/ssoui"
	"frascati/constants"
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type SsoUiService interface {
	Validate(ctx typing.Context, ticket string, callbackServ string) (string, exception.Exception)
}

type ssoUiServiceImpl struct {
	ssoClient  ssoui.Client
	authRepo   repository.AuthRepository
	jwtService auth.JwtService
}

func NewSsoUiService(ssoClient ssoui.Client, authRepo repository.AuthRepository, jwtService auth.JwtService) SsoUiService {
	return ssoUiServiceImpl{
		ssoClient:  ssoClient,
		authRepo:   authRepo,
		jwtService: jwtService,
	}
}

func (s ssoUiServiceImpl) Validate(ctx typing.Context, ticket string, callbackServ string) (string, exception.Exception) {
	ssoData, err := s.ssoClient.Validate(ticket, callbackServ)
	if err != nil {
		return "", err
	}

	userData, err := s.authRepo.FindBySsoData(ctx, ssoData.Username, ssoData.CivitasID)
	if err != nil {
		if err.Cause() != exception.CAUSE_NOT_FOUND {
			return "", err
		}

		ssoDataCopy := ssoData
		ssoDataCopy.Role = constants.ROLE_USER

		newUserData, err := s.authRepo.AddBySsoData(ctx, ssoDataCopy)
		if err != nil {
			return "", err
		}

		userData = newUserData
	}

	if !userData.HasSsoLogin {
		userData.Username = ssoData.Username
		userData.CivitasID = ssoData.CivitasID
		success, err := s.authRepo.UpdateSsoData(ctx, userData)
		if err != nil {
			return "", err
		}

		if !success {
			newErr := fmt.Errorf("failure to update sso data for registered user without prior sso login")
			return "", exception.NewBaseException(exception.CAUSE_INTERNAL, "sso/service", exception.INTERNAL, newErr)
		}
	}

	sessionData := entity.Session{
		ID:   userData.ID,
		Role: userData.Role,
	}

	token, err := s.jwtService.GenerateToken(sessionData)
	return token, err
}
