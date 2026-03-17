package queryexec_exception

import (
	"fmt"
	"frascati/exception"
)

const origin string = "queryexec"

func CreateInternalException(err error, problem string) exception.Exception {
	newErr := fmt.Errorf("%s: %w", problem, err)
	return exception.NewBaseException(exception.CAUSE_INTERNAL, origin, exception.INTERNAL, newErr)
}

func CreateRecordNotFoundException(err error, problem string) exception.Exception {
	newErr := fmt.Errorf("%s: %w", problem, err)
	return exception.NewBaseException(exception.CAUSE_NOT_FOUND, origin, exception.NOT_FOUND, newErr)
}
