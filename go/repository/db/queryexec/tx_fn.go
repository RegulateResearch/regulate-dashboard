package queryexec

import (
	"database/sql"
	"frascati/typing"
)

func getTx(ctx typing.Context) (tx *sql.Tx, ok bool) {
	val, ok := ctx.Get("tx")
	if ok {
		tx, ok = val.(*sql.Tx)
	}

	return tx, ok
}

func setTx(ctx typing.Context, tx *sql.Tx) {
	ctx.Set("tx", tx)
}
