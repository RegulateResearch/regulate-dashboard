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
	IsExistByEmailOrUsername(ctx typing.Context, email string, username string) (bool, exception.Exception)
	FindBySsoData(ctx typing.Context, username string, civitasID string) (entity.User, exception.Exception)
	AddBySsoData(ctx typing.Context, userData entity.User) (entity.User, exception.Exception)
	UpdateSsoData(ctx typing.Context, userData entity.User) (bool, exception.Exception)
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

func (r authRepositoryImpl) IsExistByEmailOrUsername(ctx typing.Context, email string, username string) (bool, exception.Exception) {
	res, err := r.repoDb.IsExistByEmailOrUsername(ctx, email, username)
	return res, err
}

func (r authRepositoryImpl) FindBySsoData(ctx typing.Context, username string, civitasID string) (entity.User, exception.Exception) {
	res, err := r.repoDb.FindBySsoData(ctx, username, civitasID)
	return res, err
}

func (r authRepositoryImpl) AddBySsoData(ctx typing.Context, userData entity.User) (entity.User, exception.Exception) {
	res, err := r.repoDb.AddBySsoData(ctx, userData)
	return res, err
}

func (r authRepositoryImpl) UpdateSsoData(ctx typing.Context, userData entity.User) (bool, exception.Exception) {
	res, err := r.repoDb.UpdateSsoData(ctx, userData)
	return res, err
}
