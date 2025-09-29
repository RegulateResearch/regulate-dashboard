package transactor

import (
	"database/sql"
	"fmt"
	"frascati/exception"
	repository_exception "frascati/repository/exception"
	"frascati/typing"
)

type Transactor interface {
	WithTransaction(ctx typing.Context, txOption Option, readOnly bool, function func(typing.Context) exception.Exception) exception.Exception
}

type transactorPostgreSQL struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) Transactor {
	return &transactorPostgreSQL{db: db}
}

func (t *transactorPostgreSQL) WithTransaction(ctx typing.Context, txOption Option, readOnly bool, fn func(typing.Context) exception.Exception) exception.Exception {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{
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
