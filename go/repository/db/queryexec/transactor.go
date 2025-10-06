package queryexec

import (
	"frascati/exception"
	"frascati/typing"
)

type Transactor interface {
	WithTransaction(ctx typing.Context, txOption TxOption, readOnly bool, function func(typing.Context) exception.Exception) exception.Exception
}
