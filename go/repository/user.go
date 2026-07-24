package repository

import (
	"frascati/exception"
	"frascati/obj/entity"
	repo_db "frascati/repository/db"
	"frascati/typing"
)

type UserRepository interface {
	FindAll(typing.Context) ([]entity.User, exception.Exception)
	FindById(ctx typing.Context, id typing.ID) (entity.User, exception.Exception)
	FilterExistingId(ctx typing.Context, ids []typing.ID) ([]entity.User, exception.Exception)
	IsExistById(ctx typing.Context, id typing.ID) (bool, exception.Exception)
	UpdateAccessBulk(ctx typing.Context, usersData []entity.User) ([]entity.User, exception.Exception)
}

type userRepositoryImpl struct {
	repoDb repo_db.UserRepository
}

func NewUserRepository(repoDb repo_db.UserRepository) UserRepository {
	return userRepositoryImpl{
		repoDb: repoDb,
	}
}

func (r userRepositoryImpl) FindAll(ctx typing.Context) ([]entity.User, exception.Exception) {
	res, err := r.repoDb.FindAll(ctx)
	return res, err
}

func (r userRepositoryImpl) FindById(ctx typing.Context, id typing.ID) (entity.User, exception.Exception) {
	res, err := r.repoDb.FindById(ctx, id)
	return res, err
}

func (r userRepositoryImpl) FilterExistingId(ctx typing.Context, ids []typing.ID) ([]entity.User, exception.Exception) {
	res, err := r.repoDb.FilterExistingId(ctx, ids)
	return res, err
}

func (r userRepositoryImpl) IsExistById(ctx typing.Context, id typing.ID) (bool, exception.Exception) {
	res, err := r.repoDb.IsExistById(ctx, id)
	return res, err
}

func (r userRepositoryImpl) UpdateAccessBulk(ctx typing.Context, usersData []entity.User) ([]entity.User, exception.Exception) {
	res, err := r.repoDb.UpdateAccessBulk(ctx, usersData)
	return res, err
}
