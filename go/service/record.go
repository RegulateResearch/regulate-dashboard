package service

import (
	"frascati/exception"
	"frascati/lambda"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
	"math/rand/v2"
)

type RecordService interface {
	FindAll(ctx typing.Context) ([]entity.Record, exception.Exception)
	AddBulk(ctx typing.Context, newData []entity.Record) ([]entity.Record, exception.Exception)
}

type recordServiceImpl struct {
	repo repository.RecordRepository
}

func NewRecordService(repo repository.RecordRepository) RecordService {
	return recordServiceImpl{
		repo: repo,
	}
}

func (s recordServiceImpl) FindAll(ctx typing.Context) ([]entity.Record, exception.Exception) {
	res, err := s.repo.FindAll(ctx)
	return res, err
}

func (s recordServiceImpl) AddBulk(ctx typing.Context, newData []entity.Record) ([]entity.Record, exception.Exception) {
	newData = lambda.MapList(newData, func(data entity.Record) entity.Record {
		return entity.Record{
			Name:        data.Name,
			Description: data.Description,
			RandNum:     rand.Int64(),
		}
	})

	res, err := s.repo.AddBulk(ctx, newData)
	return res, err
}
