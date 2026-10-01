package repository

import (
	"frascati/exception"
	"frascati/obj/entity"
	repo_db "frascati/repository/db"
	"frascati/typing"
)

type UserTaskRepository interface {
	FindByUser(ctx typing.Context, user entity.User) ([]entity.UserTask, exception.Exception)
	FindByUserAndId(ctx typing.Context, user entity.User, ID typing.ID) (entity.UserTask, exception.Exception)
	AddInferByNotYetAdded(ctx typing.Context, user entity.User) (dataAffected int64, err exception.Exception)
	UpdateStatusAndTimes(ctx typing.Context, task entity.UserTask) (entity.UserTask, exception.Exception)
}

type userTaskRepositoryImpl struct {
	repoDb repo_db.UserTaskRepository
}

func NewUserTaskRepository(repoDb repo_db.UserTaskRepository) UserTaskRepository {
	return userTaskRepositoryImpl{
		repoDb: repoDb,
	}
}

func (r userTaskRepositoryImpl) FindByUser(ctx typing.Context, user entity.User) ([]entity.UserTask, exception.Exception) {
	res, err := r.repoDb.FindByUser(ctx, user)
	return res, err
}

func (r userTaskRepositoryImpl) FindByUserAndId(ctx typing.Context, user entity.User, id typing.ID) (entity.UserTask, exception.Exception) {
	res, err := r.repoDb.FindByUserAndId(ctx, user, id)
	return res, err
}

func (r userTaskRepositoryImpl) AddInferByNotYetAdded(ctx typing.Context, user entity.User) (dataAffected int64, err exception.Exception) {
	res, err := r.repoDb.AddInferByNotYetAdded(ctx, user)
	return res, err
}

func (r userTaskRepositoryImpl) UpdateStatusAndTimes(ctx typing.Context, task entity.UserTask) (entity.UserTask, exception.Exception) {
	res, err := r.repoDb.UpdateStatusAndTimes(ctx, task)
	return res, err
}
