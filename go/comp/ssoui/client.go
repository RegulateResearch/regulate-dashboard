package ssoui

import (
	"fmt"
	"frascati/exception"
	"frascati/obj/entity"
	xmlutils "frascati/utils/xml"
	"io"
	"net/http"
)

type Client interface {
	Validate(ticket string, callbackServ string) (entity.User, exception.Exception)
}

type client struct {
	url string
}

func NewSsoClient(ssoUrl string) Client {
	return client{
		url: ssoUrl,
	}
}

func (c client) Validate(ticket string, callbackServ string) (entity.User, exception.Exception) {
	link := fmt.Sprintf("%s/serviceValidate?ticket=%s&service=%s", c.url, ticket, callbackServ)
	resp, err := http.Get(link)
	if err != nil {
		newErr := fmt.Errorf("cannot access SSO UI: %w", err)
		return entity.User{}, exception.NewBaseException(exception.CAUSE_INTERNAL, "ssoui", exception.INTERNAL, newErr)
	}

	body := resp.Body
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		newErr := fmt.Errorf("cannot read SSO response: %w", err)
		return entity.User{}, exception.NewBaseException(exception.CAUSE_INTERNAL, "ssoui", exception.INTERNAL, newErr)
	}

	res, parseErr := c.parse(data)

	username := res.AuthSuccess.Username
	fullname := res.AuthSuccess.Attr.Name
	npm := res.AuthSuccess.Attr.Npm

	userData := entity.User{
		Username:    username,
		DisplayName: fullname,
		CivitasID:   npm,
	}

	return userData, parseErr
}

func (c client) parse(data []byte) (successResponse, exception.Exception) {
	res, err := xmlutils.ParseGeneric(data, newSuccessResponse)
	if err != nil {
		failRes, failParseErr := xmlutils.ParseGeneric(data, newFailureResponse)
		if failParseErr != nil {
			return successResponse{}, failParseErr
		}

		failCode := failRes.AuthFailure.Code
		failMessage := failRes.AuthFailure.Message
		newErr := fmt.Errorf("SSO UI validation fails; code = %s, message = %s", failCode, failMessage)

		return successResponse{}, exception.NewBaseException(exception.CAUSE_UNAUTHORIZED, "ssoui", "authentication fails", newErr)
	}

	return res, nil
}
