package queryexec

import (
	"database/sql"
	"frascati/typing"
)

type QueryExecutor interface {
	QueryContext(ctx typing.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx typing.Context, query string, args ...any) *sql.Row
	ExecContext(ctx typing.Context, query string, args ...any) (sql.Result, error)
}
