package entity

type Record struct {
	Base
	Name        string
	RandNum     int64
	Description string
}

func NewRecord() Record {
	return Record{
		Base: newBase(),
	}
}
