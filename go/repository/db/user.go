package repo_db

import (
	"frascati/comp/queryexec"
	"frascati/exception"
	"frascati/obj/entity"
	repository_exception "frascati/repository/exception"
	"frascati/typing"
	"frascati/utils/querying"
)

type UserRepository interface {
	FindAll(typing.Context) ([]entity.User, exception.Exception)
	FindById(ctx typing.Context, id typing.ID) (entity.User, exception.Exception)
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
	query :=
		`SELECT id, email, username, user_role
		FROM users`

	rows, err := r.executor.QueryContext(ctx, query)
	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user")
	}
	defer r.executor.CloseRows(rows, "user - FindAll")

	res, err := querying.ScanForRows(
		rows, entity.NewUser,
		func(rows queryexec.Rows, elem entity.User) (entity.User, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.Email, &elem.Username, &elem.Role)
			return elem, err
		},
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user")
	}

	return res, nil
}

func (r userRepositoryImpl) FindById(ctx typing.Context, id typing.ID) (entity.User, exception.Exception) {
	var res entity.User
	querystr := `
		SELECT id, email, username, user_role
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`

	err := r.executor.QueryRowContext(ctx, querystr, id).Scan(&res.ID, &res.Email, &res.Username, &res.Role)
	if err != nil {
		return entity.User{}, repository_exception.WrapQueryexecException(err, "user")
	}

	return res, nil
}
