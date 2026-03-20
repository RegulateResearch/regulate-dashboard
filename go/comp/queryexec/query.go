package queryexec

import (
	"frascati/exception"
	"frascati/typing"
)

type QueryExecutor interface {
	QueryContext(ctx typing.Context, query string, args ...any) (Rows, exception.Exception)
	QueryRowContext(ctx typing.Context, query string, args ...any) Row
	ExecContext(ctx typing.Context, query string, args ...any) (Result, exception.Exception)
	CloseRows(rows Rows, identifier string)
}
