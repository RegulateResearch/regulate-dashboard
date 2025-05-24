package repository_exception

import (
	"fmt"
	"frascati/exception"
)

func CreateDBException(err error, domain string, message string) exception.Exception {
	return exception.NewBaseException(exception.CAUSE_INTERNAL, fmt.Sprintf("%s/repository", domain), message, err)
}

func CreateRecordNotFoundException(err error, domain string, message string) exception.Exception {
	return exception.NewBaseException(exception.CAUSE_NOT_FOUND, fmt.Sprintf("%s/repository", domain), message, err)
}
