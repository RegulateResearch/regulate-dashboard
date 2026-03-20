package txhandler

type TxOption int

const (
	TxOptionDefault TxOption = iota
	TxOptionReadCommitted
	TxOptionWriteCommitted
	TxOptionRepeatableRead
	TxOptionSerializable
)
