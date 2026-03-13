package dbhandler

import (
	"database/sql"
	"frascati/comp/txhandler"
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

func transformTxOptionToIsolationLevel(opt txhandler.TxOption) sql.IsolationLevel {
	res := sql.LevelDefault
	optLevelMap := map[txhandler.TxOption]sql.IsolationLevel{
		txhandler.TxOptionDefault:        sql.LevelDefault,
		txhandler.TxOptionReadCommitted:  sql.LevelReadCommitted,
		txhandler.TxOptionRepeatableRead: sql.LevelRepeatableRead,
		txhandler.TxOptionWriteCommitted: sql.LevelWriteCommitted,
		txhandler.TxOptionSerializable:   sql.LevelSerializable,
	}

	level, ok := optLevelMap[opt]
	if ok {
		res = level
	}

	return res
}
