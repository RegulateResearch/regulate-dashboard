package repo_db

import (
	"fmt"
	"frascati/comp/background"
	"frascati/comp/queryexec"
	"frascati/exception"
)

type baseDbRepository interface {
	getExecutor() queryexec.QueryExecutor
	closeRows(rows queryexec.Rows, identifier string)
}

type baseDbRepositoryImpl struct {
	processor background.Processor
	executor  queryexec.QueryExecutor
}

func newBaseDbRepository(executor queryexec.QueryExecutor, processor background.Processor) baseDbRepository {
	return baseDbRepositoryImpl{
		processor: processor,
		executor:  executor,
	}
}

func (r baseDbRepositoryImpl) getExecutor() queryexec.QueryExecutor {
	return r.executor
}

func (r baseDbRepositoryImpl) closeRows(rows queryexec.Rows, identifier string) {
	err := rows.Close()
	if err != nil {
		taskName := fmt.Sprintf("%s - row close", identifier)
		r.processor.AddTask(taskName, func() (any, exception.Exception) {
			return nil, err
		})
	}
}
