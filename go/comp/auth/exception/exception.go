package auth_exception

import (
	"errors"
	"fmt"
	"frascati/exception"
)

var errEmailAlreadyExists = errors.New("email is already registered")

func GenerateErrPasswordFailure(err error) exception.Exception {
	return exception.NewBaseException(exception.CAUSE_INTERNAL, "auth/bcrypt", "bcrypt hashing error", err)
}

func GenerateErrUserAlreadyExist() exception.Exception {
	return exception.NewBaseException(exception.CAUSE_USER, "auth/service", "email already used", errEmailAlreadyExists)
}

func GenerateErrLoginFail(err error) exception.Exception {
	return exception.NewBaseException(exception.CAUSE_UNAUTHORIZED, "auth/service", "either email or password is wrong", err)
}

func GenerateErrAuthFailComposite(err exception.Exception) exception.Exception {
	return exception.NewCompositeException(exception.CAUSE_INTERNAL, "auth/service", "something is wrong in our end", err)
}

func GenerateErrAuthFailBase(subdomain string, err error) exception.Exception {
	return exception.NewBaseException(exception.CAUSE_INTERNAL, fmt.Sprintf("auth/%s", subdomain), "auth process fail", err)
}

func GenerateErrInvalidToken(subdomain string, err error) exception.Exception {
	return exception.NewBaseException(exception.CAUSE_UNAUTHORIZED, fmt.Sprintf("auth/%s", subdomain), "token is invalid", err)
}
