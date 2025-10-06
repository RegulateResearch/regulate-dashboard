package queryexec

import "database/sql"

type TxOption int

const (
	TxOptionDefault TxOption = iota
	TxOptionReadCommitted
	TxOptionWriteCommitted
	TxOptionRepeatableRead
	TxOptionSerializable
)

func (o TxOption) toSqlIsolationLevel() sql.IsolationLevel {
	allOptions := []sql.IsolationLevel{
		sql.LevelDefault,
		sql.LevelReadCommitted,
		sql.LevelWriteCommitted,
		sql.LevelRepeatableRead,
		sql.LevelSerializable,
	}

	optionNum := int(o)
	selectedOption := sql.LevelDefault
	if optionNum >= 0 && optionNum < len(allOptions) {
		selectedOption = allOptions[optionNum]
	}

	return selectedOption
}
