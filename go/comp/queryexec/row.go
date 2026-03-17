package queryexec

import (
	"database/sql"
	"errors"
	queryexec_exception "frascati/comp/queryexec/exception"
	"frascati/exception"
)

type Row interface {
	Scan(dest ...any) exception.Exception
}

type row struct {
	sqlrow *sql.Row
}

func NewRow(sqlrow *sql.Row) Row {
	return row{
		sqlrow: sqlrow,
	}
}

func (r row) Scan(dest ...any) exception.Exception {
	err := r.sqlrow.Scan(dest...)
	if err != nil {
		problem := "single row scan error"
		if errors.Is(err, sql.ErrNoRows) {
			return queryexec_exception.CreateRecordNotFoundException(err, problem)
		}

		return queryexec_exception.CreateInternalException(err, problem)
	}

	return nil
}
