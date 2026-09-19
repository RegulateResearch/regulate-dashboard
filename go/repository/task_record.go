package repository

import repo_db "frascati/repository/db"

type TaskRecordRepository interface {
}

type taskRecordRepositoryImpl struct {
	repo repo_db.TaskRecordRepository
}

func NewTaskRecordRepository(repo repo_db.TaskRecordRepository) TaskRecordRepository {
	return taskRecordRepositoryImpl{
		repo: repo,
	}
}
