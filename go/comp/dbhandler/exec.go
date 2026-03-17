package dbhandler

import (
	"context"
	"database/sql"
	"fmt"
	"frascati/comp/background"
	"frascati/comp/queryexec"
	queryexec_exception "frascati/comp/queryexec/exception"
	"frascati/comp/txhandler"
	tx_exception "frascati/comp/txhandler/exception"
	"frascati/exception"
	"frascati/typing"
)

type DbExecutor interface {
	queryexec.QueryExecutor
	txhandler.Transactor
}

type dbExecutorImpl struct {
	db        *sql.DB
	processor background.Processor
}

func NewDbExecutor(db *sql.DB, processor background.Processor) DbExecutor {
	return dbExecutorImpl{
		db:        db,
		processor: processor,
	}
}

func (e dbExecutorImpl) selectRunner(ctx typing.Context) queryRunner {
	var runner queryRunner = e.db
	tx, ok := getTx(ctx)
	if ok {
		runner = tx
	}

	return runner
}

func (e dbExecutorImpl) QueryContext(ctx typing.Context, query string, args ...any) (queryexec.Rows, exception.Exception) {
	runner := e.selectRunner(ctx)
	rows, err := runner.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, queryexec_exception.CreateInternalException(err, "query error")
	}

	return queryexec.NewRows(rows), nil
}

func (e dbExecutorImpl) QueryRowContext(ctx typing.Context, query string, args ...any) queryexec.Row {
	runner := e.selectRunner(ctx)
	row := runner.QueryRowContext(ctx, query, args...)

	return queryexec.NewRow(row)
}

func (e dbExecutorImpl) ExecContext(ctx typing.Context, query string, args ...any) (queryexec.Result, exception.Exception) {
	runner := e.selectRunner(ctx)
	result, err := runner.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, queryexec_exception.CreateInternalException(err, "query exec error")
	}

	return queryexec.NewResult(result), nil
}

func (e dbExecutorImpl) CloseRows(rows queryexec.Rows, identifier string) {
	err := rows.Close()
	if err != nil {
		taskName := fmt.Sprintf("%s - row close", identifier)
		e.processor.AddTask(taskName, func() (any, exception.Exception) {
			return nil, err
		})
	}
}

func (q dbExecutorImpl) WithTransaction(ctx typing.Context, txOption txhandler.TxOption, readOnly bool, fn func(typing.Context) exception.Exception) exception.Exception {
	tx, err := q.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: transformTxOptionToIsolationLevel(txOption),
		ReadOnly:  readOnly,
	})

	if err != nil {
		return tx_exception.TransactionError(err, "cannot begin transaction")
	}
	defer tx.Rollback()

	setTx(ctx, tx)

	exc := fn(ctx)
	if exc != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return tx_exception.TransactionError(errRollback, "cannot rollback transaction")
		}
		return exc
	}

	if err := tx.Commit(); err != nil {
		return tx_exception.TransactionError(err, "cannot commit transaction")
	}

	return nil
}

type queryRunner interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}
