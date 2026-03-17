package queryexec

import (
	"database/sql"
	queryexec_exception "frascati/comp/queryexec/exception"
	"frascati/exception"
)

type Result interface {
	RowsAffected() (int64, exception.Exception)
}

type result struct {
	res sql.Result
}

func NewResult(res sql.Result) Result {
	return result{
		res: res,
	}
}

func (r result) RowsAffected() (int64, exception.Exception) {
	rowsAffected, err := r.res.RowsAffected()
	if err != nil {
		return 0, queryexec_exception.CreateInternalException(err, "rows affected error")
	}

	return rowsAffected, nil
}
