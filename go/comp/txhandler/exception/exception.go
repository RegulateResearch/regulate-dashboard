package tx_exception

import "frascati/exception"

func TransactionError(err error) exception.Exception {
	return exception.NewBaseException(exception.CAUSE_INTERNAL, "tx", exception.INTERNAL, err)
}
