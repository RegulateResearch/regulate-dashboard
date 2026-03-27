package repository

import (
	"frascati/exception"
	"frascati/obj/entity"
	repo_db "frascati/repository/db"
	"frascati/typing"
)

type RecordRepository interface {
	FindAll(ctx typing.Context) ([]entity.Record, exception.Exception)
	AddBulk(ctx typing.Context, newData []entity.Record) ([]entity.Record, exception.Exception)
}

type recordRepositoryImpl struct {
	repoDb repo_db.RecordRepository
}

func NewRecordRepository(repoDb repo_db.RecordRepository) RecordRepository {
	return recordRepositoryImpl{
		repoDb: repoDb,
	}
}

func (r recordRepositoryImpl) FindAll(ctx typing.Context) ([]entity.Record, exception.Exception) {
	res, err := r.repoDb.FindAll(ctx)
	return res, err
}

func (r recordRepositoryImpl) AddBulk(ctx typing.Context, newData []entity.Record) ([]entity.Record, exception.Exception) {
	res, err := r.repoDb.AddBulk(ctx, newData)
	return res, err
}
