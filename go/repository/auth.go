package repository

import (
	"frascati/exception"
	"frascati/obj/entity"
	repo_db "frascati/repository/db"
	"frascati/typing"
)

type AuthRepository interface {
	Add(ctx typing.Context, newUserData entity.User) (entity.User, exception.Exception)
	FindByEmail(ctx typing.Context, email string) (entity.User, exception.Exception)
	IsExistByEmail(ctx typing.Context, email string) (bool, exception.Exception)
}

type authRepositoryImpl struct {
	repoDb repo_db.AuthRepository
}

func NewAuthRepository(repoDb repo_db.AuthRepository) AuthRepository {
	return authRepositoryImpl{repoDb: repoDb}
}

func (r authRepositoryImpl) Add(ctx typing.Context, newUserData entity.User) (entity.User, exception.Exception) {
	res, err := r.repoDb.Add(ctx, newUserData)
	return res, err
}

func (r authRepositoryImpl) FindByEmail(ctx typing.Context, email string) (entity.User, exception.Exception) {
	res, err := r.repoDb.FindByEmail(ctx, email)
	return res, err
}

func (r authRepositoryImpl) IsExistByEmail(ctx typing.Context, email string) (bool, exception.Exception) {
	res, err := r.repoDb.IsExistByEmail(ctx, email)
	return res, err
}
