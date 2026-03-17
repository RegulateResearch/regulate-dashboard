package repo_db

import (
	"frascati/comp/queryexec"
	"frascati/exception"
	"frascati/obj/entity"
	repository_exception "frascati/repository/exception"
	"frascati/typing"
)

type UserRepository interface {
	FindAll(typing.Context) ([]entity.User, exception.Exception)
}

type userRepositoryImpl struct {
	executor queryexec.QueryExecutor
}

func NewUserRepository(executor queryexec.QueryExecutor) UserRepository {
	return userRepositoryImpl{
		executor: executor,
	}
}

func (r userRepositoryImpl) FindAll(ctx typing.Context) ([]entity.User, exception.Exception) {
	res := make([]entity.User, 0)
	query :=
		`SELECT id, username, user_role
		FROM users`

	rows, err := r.executor.QueryContext(ctx, query)
	if err != nil {
		return nil, repository_exception.CreateDBException(err, "user", "something is wrong in our end")
	}
	defer r.executor.CloseRows(rows, "user - FindAll")

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Username, &user.Role)
		if err != nil {
			return nil, repository_exception.WrapQueryexecException(err, "user")
		}

		res = append(res, user)
	}

	return res, nil
}
