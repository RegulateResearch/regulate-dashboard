package queryexec

import (
	"context"
	"database/sql"
	"fmt"
	"frascati/exception"
	repository_exception "frascati/repository/exception"
	"frascati/typing"
)

type DbExecutor interface {
	QueryExecutor
	Transactor
}

type dbExecutorImpl struct {
	db *sql.DB
}

func NewQueryExecutor(db *sql.DB) DbExecutor {
	return dbExecutorImpl{db: db}
}

func (e dbExecutorImpl) selectRunner(ctx typing.Context) queryRunner {
	var runner queryRunner = e.db
	tx, ok := getTx(ctx)
	if ok {
		runner = tx
	}

	return runner
}

func (e dbExecutorImpl) QueryContext(ctx typing.Context, query string, args ...any) (*sql.Rows, error) {
	runner := e.selectRunner(ctx)
	rows, err := runner.QueryContext(ctx, query, args...)

	return rows, err
}

func (e dbExecutorImpl) QueryRowContext(ctx typing.Context, query string, args ...any) *sql.Row {
	runner := e.selectRunner(ctx)
	row := runner.QueryRowContext(ctx, query, args...)

	return row
}

func (e dbExecutorImpl) ExecContext(ctx typing.Context, query string, args ...any) (sql.Result, error) {
	runner := e.selectRunner(ctx)
	result, err := runner.ExecContext(ctx, query, args...)

	return result, err
}

func (q dbExecutorImpl) WithTransaction(ctx typing.Context, txOption TxOption, readOnly bool, fn func(typing.Context) exception.Exception) exception.Exception {
	tx, err := q.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: txOption.toSqlIsolationLevel(),
		ReadOnly:  readOnly,
	})
	if err != nil {
		return repository_exception.CreateDBException(err, "transactor", "cannot begin transaction")
	}
	defer tx.Rollback()

	setTx(ctx, tx)

	exc := fn(ctx)
	if exc != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return repository_exception.CreateDBException(errRollback, "transactor", fmt.Sprintf("cannot rollback transaction: %s", exc.Error()))
		}
		return exc
	}

	if err := tx.Commit(); err != nil {
		return repository_exception.CreateDBException(err, "transactor", "cannot commit transaction")
	}

	return nil
}

type queryRunner interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}
