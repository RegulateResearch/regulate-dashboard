package repo_db

import "frascati/comp/queryexec"

type TaskRecordRepository interface {
}

type taskRecordRepositoryImpl struct {
	executor queryexec.QueryExecutor
}

func NewTaskRecordRepository(executor queryexec.QueryExecutor) TaskRecordRepository {
	return taskRecordRepositoryImpl{
		executor: executor,
	}
}
