package dao

// only for scanning
// should not be used outside repo_db
type RecordDb struct {
	BaseDb
	Name        string
	RandNum     int64
	Description string
}

func NewRecordDb() RecordDb {
	return RecordDb{
		BaseDb: newBaseDb(),
	}
}
