package exception

import (
	"fmt"
	"frascati/exception"
)

func GenerateErrPasswordFailure(err error) exception.Exception {
	return exception.NewBaseException(exception.CAUSE_INTERNAL, "auth/bcrypt", "bcrypt hashing error", err)
}

func GenerateErrUserAlreadyExist() exception.Exception {
	return exception.NewBaseException(exception.CAUSE_USER, "auth/service", "username already used", nil)
}

func GenerateErrLoginFail(err error) exception.Exception {
	return exception.NewBaseException(exception.CAUSE_UNAUTHORIZED, "auth/service", "either username or password is wrong", err)
}

func GenerateErrAuthFailComposite(err exception.Exception) exception.Exception {
	return exception.NewCompositeException(exception.CAUSE_INTERNAL, "auth/service", "something is wrong in our end", err)
}

func GenerateErrAuthFailBase(subdomain string, err error) exception.Exception {
	return exception.NewBaseException(exception.CAUSE_INTERNAL, fmt.Sprintf("auth/%s", subdomain), "auth process fail", err)
}
