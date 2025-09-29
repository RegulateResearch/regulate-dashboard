package transactor

import (
	"context"
	"database/sql"
	"fmt"
	"frascati/exception"
	repository_exception "frascati/repository/exception"
)

type Transactor interface {
	WithTransaction(ctx context.Context, txOption Option, readOnly bool, function func(context.Context) exception.Exception) exception.Exception
}

type transactorPostgreSQL struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) Transactor {
	return &transactorPostgreSQL{db: db}
}

func ExtractTx(ctx context.Context) *sql.Tx {
	tx, ok := ctx.Value(TxKey{}).(*sql.Tx)
	if !ok {
		return nil
	}

	return tx
}

type TxKey struct{}

func (t *transactorPostgreSQL) WithTransaction(ctx context.Context, txOption Option, readOnly bool, fn func(ctx context.Context) exception.Exception) exception.Exception {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: txOption.toSqlIsolationLevel(),
		ReadOnly:  readOnly,
	})
	if err != nil {
		return repository_exception.CreateDBException(err, "transactor", "cannot begin transaction")
	}
	defer tx.Rollback()

	exc := fn(context.WithValue(ctx, TxKey{}, tx))
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
