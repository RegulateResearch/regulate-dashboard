package transactor

import "database/sql"

type Option int

const (
	OptionDefault Option = iota
	OptionReadCommitted
	OptionWriteCommitted
	OptionRepeatableRead
	OptionSerializable
)

func (o Option) toSqlIsolationLevel() sql.IsolationLevel {
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
