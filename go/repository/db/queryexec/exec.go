package queryexec

import (
	"context"
	"database/sql"
	"frascati/repository/transactor"
	"frascati/typing"
)

type queryRunner interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type QueryExecutor interface {
	QueryContext(ctx typing.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx typing.Context, query string, args ...any) *sql.Row
	ExecContext(ctx typing.Context, query string, args ...any) (sql.Result, error)
}

type queryExecutorImpl struct {
	db *sql.DB
}

func NewQueryExecutor(db *sql.DB) QueryExecutor {
	return queryExecutorImpl{db: db}
}

func (e queryExecutorImpl) selectRunner(ctx typing.Context) queryRunner {
	var runner queryRunner = e.db
	tx, ok := transactor.GetTx(ctx)
	if ok {
		runner = tx
	}

	return runner
}

func (e queryExecutorImpl) QueryContext(ctx typing.Context, query string, args ...any) (*sql.Rows, error) {
	runner := e.selectRunner(ctx)
	rows, err := runner.QueryContext(ctx, query, args...)

	return rows, err
}

func (e queryExecutorImpl) QueryRowContext(ctx typing.Context, query string, args ...any) *sql.Row {
	runner := e.selectRunner(ctx)
	row := runner.QueryRowContext(ctx, query, args...)

	return row
}

func (e queryExecutorImpl) ExecContext(ctx typing.Context, query string, args ...any) (sql.Result, error) {
	runner := e.selectRunner(ctx)
	result, err := runner.ExecContext(ctx, query, args...)

	return result, err
}
