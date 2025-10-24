package repository

import (
	"frascati/exception"
	"frascati/obj/entity"
	repo_db "frascati/repository/db"
	"frascati/typing"
)

type UserRepository interface {
	FindAll(typing.Context) ([]entity.User, exception.Exception)
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
