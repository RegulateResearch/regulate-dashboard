package repo_db

import (
	"fmt"
	"frascati/comp/queryexec"
	"frascati/exception"
	"frascati/lambda"
	"frascati/obj/converter"
	"frascati/obj/dao"
	"frascati/obj/entity"
	repository_exception "frascati/repository/exception"
	"frascati/typing"
	"frascati/utils/querying"
	"log"
)

type UserRepository interface {
	FindAll(typing.Context) ([]entity.User, exception.Exception)
	FindById(ctx typing.Context, id typing.ID) (entity.User, exception.Exception)
	FilterExistingId(ctx typing.Context, ids []typing.ID) ([]entity.User, exception.Exception)
	IsExistById(ctx typing.Context, id typing.ID) (bool, exception.Exception)
	UpdateAccessBulk(ctx typing.Context, usersData []entity.User) ([]entity.User, exception.Exception)
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
		`SELECT id, email, username, display_name, user_role, civitas_id, academic_role
		FROM users`

	rows, err := r.executor.QueryContext(ctx, query)
	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user")
	}
	defer r.executor.CloseRows(rows, "user - FindAll")

	res, err := querying.ScanForRowsThenTransform(
		rows, dao.NewUserDb,
		func(rows queryexec.Rows, elem dao.UserDb) (dao.UserDb, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.Email, &elem.Username, &elem.DisplayName, &elem.Role, &elem.CivitasID, &elem.AcademicRole)
			return elem, err
		},
		converter.UserDbToEntity,
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user")
	}

	return res, nil
}

func (r userRepositoryImpl) FindById(ctx typing.Context, id typing.ID) (entity.User, exception.Exception) {
	var res dao.UserDb
	querystr := `
		SELECT 
			id, email, username, display_name, user_role, civitas_id
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`

	err := r.executor.QueryRowContext(ctx, querystr, id).
		Scan(&res.ID, &res.Email, &res.Username, &res.DisplayName, &res.Role, &res.CivitasID)
	if err != nil {
		return entity.User{}, repository_exception.WrapQueryexecException(err, "user")
	}

	return converter.UserDbToEntity(res), nil
}

func (r userRepositoryImpl) FilterExistingId(ctx typing.Context, ids []typing.ID) ([]entity.User, exception.Exception) {
	querystr := `
		SELECT DISTINCT u.id, u.academic_role
		FROM (
			VALUES
				%s
		) AS ids(id)
		JOIN (
			SELECT id, academic_role
			FROM users
			WHERE deleted_at IS NULL
		) AS u ON ids.id = u.id
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
	rows, err := r.executor.QueryContext(ctx, querystr, args...)
	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user")
	}
	defer r.executor.CloseRows(rows, "user - FilterExistingId")

	resDb, err := querying.ScanForRows(
		rows, dao.NewUserDb,
		func(rows queryexec.Rows, elem dao.UserDb) (dao.UserDb, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.AcademicRole)
			return elem, err
		},
	)

	res := lambda.MapList(resDb, converter.UserDbToEntity)
	resDto := lambda.MapList(res, converter.UserEntityToDTO)
	log.Println(resDto)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user")
	}

	return res, nil
}

func (r userRepositoryImpl) IsExistById(ctx typing.Context, id typing.ID) (bool, exception.Exception) {
	querystr := `
		SELECT 1
		FROM users
		WHERE
			id = $1 AND
			deleted_at IS NULL
	`

	res, err := r.executor.ExecContext(ctx, querystr, id)
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "user")
	}

	rowsCount, err := res.RowsAffected()
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "user")
	}

	return rowsCount > 0, nil
}

func (r userRepositoryImpl) UpdateAccessBulk(ctx typing.Context, usersData []entity.User) ([]entity.User, exception.Exception) {
	querystr := `
		UPDATE users
		SET
			user_role = update_data.role,
			academic_role = update_data.academic_role,
			updated_at = NOW()
		FROM (
			VALUES
				%s
		) AS update_data(id, role, academic_role)
		WHERE
			users.id = update_data.id AND
			users.deleted_at IS NULL
		RETURNING users.id, users.user_role, users.academic_role 
	`
	rowstr, args, _ := querying.BulkValuesFromBeginning(usersData, func(user entity.User, currentIdx int) querying.DataStrArgs {
		return querying.DataStrArgs{
			RowStr: fmt.Sprintf(
				"(CAST($%d AS BIGINT), CAST($%d AS SMALLINT), CAST($%d AS SMALLINT))",
				currentIdx, currentIdx+1, currentIdx+2,
			),
			Args: []any{user.ID, user.Role, user.AcademicRole},
		}
	})

	querystr = fmt.Sprintf(querystr, rowstr)
	log.Println(args...)
	rows, err := r.executor.QueryContext(ctx, querystr, args...)

	res, err := querying.ScanForRowsThenTransform(
		rows, dao.NewUserDb,
		func(rows queryexec.Rows, elem dao.UserDb) (dao.UserDb, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.Role, &elem.AcademicRole)
			return elem, err
		},
		converter.UserDbToEntity,
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user")
	}

	return res, nil

}
