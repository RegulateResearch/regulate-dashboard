package tx_exception

import "frascati/exception"

func TransactionError(err error, message string) exception.Exception {
	return nil
}
