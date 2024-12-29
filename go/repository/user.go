package repository

import (
	"context"
	"database/sql"
	"frascati/entity"
	"frascati/exception"
	repository_exception "frascati/repository/exception"
	"log"
)

type UserRepository interface {
	FindAll(context.Context) ([]entity.User, exception.Exception)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return userRepositoryImpl{
		db: db,
	}
}

func (r userRepositoryImpl) FindAll(ctx context.Context) ([]entity.User, exception.Exception) {
	res := make([]entity.User, 0)
	query :=
		`SELECT id, username, user_role
		FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, repository_exception.CreateDBException(err, "user", "something is wrong in our end")
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println("cannot close rows, ", err.Error())
		}
	}()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Username, &user.Role)
		if err != nil {
			return nil, repository_exception.CreateDBException(err, "user", "something is wrong in our end")
		}

		res = append(res, user)
	}

	return res, nil
}
