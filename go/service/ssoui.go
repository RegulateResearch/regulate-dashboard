package service

import (
	"frascati/comp/ssoui"
	"frascati/exception"
)

type SsoUiService interface {
	Validate(ticket string, callbackServ string) (any, exception.Exception)
}

type ssoUiServiceImpl struct {
	ssoClient ssoui.Client
}

func NewSsoUiService(ssoClient ssoui.Client) SsoUiService {
	return ssoUiServiceImpl{
		ssoClient: ssoClient,
	}
}

func (s ssoUiServiceImpl) Validate(ticket string, callbackServ string) (any, exception.Exception) {
	res, err := s.ssoClient.Validate(ticket, callbackServ)
	return res, err
}
