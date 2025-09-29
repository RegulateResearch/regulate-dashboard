package repo_db

import (
	"database/sql"
	"errors"
	"frascati/constants"
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository/db/queryexec"
	repository_exception "frascati/repository/exception"
	"frascati/typing"
)

type AuthRepositoryDb interface {
	Add(ctx typing.Context, newUserData entity.User) (entity.User, exception.Exception)
	FindByEmail(ctx typing.Context, email string) (entity.User, exception.Exception)
	IsExistByEmail(ctx typing.Context, email string) (bool, exception.Exception)
}

type authRepositoryDbImpl struct {
	executor queryexec.QueryExecutor
}

func NewAuthRepositoryDb(executor queryexec.QueryExecutor) AuthRepositoryDb {
	return authRepositoryDbImpl{executor: executor}
}

func (r authRepositoryDbImpl) Add(ctx typing.Context, newUserData entity.User) (entity.User, exception.Exception) {
	query := `
		INSERT INTO users(email, username, password, user_role, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, email, username, user_role
	`

	var user entity.User
	err := r.executor.QueryRowContext(ctx, query,
		newUserData.Email, newUserData.Username, newUserData.Password, constants.ROLE_USER).
		Scan(&user.ID, &user.Email, &user.Username, &user.Role)
	if err != nil {
		return entity.User{}, repository_exception.CreateDBException(err, "auth", "something is wrong in our end")
	}

	return user, nil
}

func (r authRepositoryDbImpl) FindByEmail(ctx typing.Context, email string) (entity.User, exception.Exception) {
	query := `
		SELECT id, username, password, user_role
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	var user entity.User
	err := r.executor.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		var empty entity.User
		if errors.Is(err, sql.ErrNoRows) {
			return empty, repository_exception.CreateRecordNotFoundException(err, "auth", "record not found")
		}

		return empty, repository_exception.CreateDBException(err, "auth", "something is wrong in our end")
	}

	user.Email = email

	return user, nil
}

func (r authRepositoryDbImpl) IsExistByEmail(ctx typing.Context, email string) (bool, exception.Exception) {
	query := `SELECT 1 FROM users WHERE email = $1`
	res, err := r.executor.ExecContext(ctx, query, email)
	if err != nil {
		return false, repository_exception.CreateDBException(err, "auth", "something is wrong in our end")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, repository_exception.CreateDBException(err, "auth", "something is wrong in our end")
	}

	return rowsAffected > 0, nil
}
