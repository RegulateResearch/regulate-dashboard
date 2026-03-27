package repo_db

import (
	"fmt"
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
	FilterExistingId(ctx typing.Context, ids []typing.ID) ([]typing.ID, exception.Exception)
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

func (r userRepositoryImpl) FilterExistingId(ctx typing.Context, ids []typing.ID) ([]typing.ID, exception.Exception) {
	querystr := `
		SELECT id
		FROM (
			VALUES
				%s
		) AS ids(id)
		WHERE EXISTS (
			SELECT 1
			FROM users AS u
			WHERE u.id = ids.id AND deleted_at IS NULL
		)
	`

	paramIdxStart := 1
	rowstr, args, _ := querying.BulkValues(
		ids, paramIdxStart,
		func(id typing.ID, currentIdx int) querying.DataStrArgs {
			return querying.DataStrArgs{
				RowStr: fmt.Sprintf("(CAST($%d AS BIGINT))", currentIdx),
				Args:   []any{id},
			}
		},
	)

	querystr = fmt.Sprintf(querystr, rowstr)
	fmt.Println(querystr)
	rows, err := r.executor.QueryContext(ctx, querystr, args...)
	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user")
	}
	defer r.executor.CloseRows(rows, "user - FilterExistingId")

	res, err := querying.ScanForRows(
		rows, typing.IDDefault,
		func(rows queryexec.Rows, elem typing.ID) (typing.ID, exception.Exception) {
			err := rows.Scan(&elem)
			return elem, err
		},
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user")
	}

	return res, nil
}
