package queryexec

import (
	"database/sql"
	queryexec_exception "frascati/comp/queryexec/exception"
	"frascati/exception"
)

type Rows interface {
	Scan(dest ...any) exception.Exception
	Next() bool
	Close() exception.Exception
}

type rows struct {
	sqlrows *sql.Rows
}

func NewRows(sqlrows *sql.Rows) Rows {
	return rows{
		sqlrows: sqlrows,
	}
}

func (r rows) Scan(dest ...any) exception.Exception {
	err := r.sqlrows.Scan(dest...)
	if err != nil {
		return queryexec_exception.CreateInternalException(err, "multiple row scan error")
	}

	return nil
}

func (r rows) Next() bool {
	return r.sqlrows.Next()
}

func (r rows) Close() exception.Exception {
	err := r.sqlrows.Close()
	if err != nil {
		return queryexec_exception.CreateInternalException(err, "row close error")
	}

	return nil
}
