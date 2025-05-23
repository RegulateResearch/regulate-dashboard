package repository

import (
	"context"
	"database/sql"
	"errors"
	"frascati/constants"
	"frascati/entity"
	"frascati/exception"
	repository_exception "frascati/repository/exception"
)

type AuthRepository interface {
	Add(ctx context.Context, newUserData entity.UserWrite) (entity.User, exception.Exception)
	FindByUsername(ctx context.Context, username string) (entity.User, exception.Exception)
	IsExist(ctx context.Context, username string) (bool, exception.Exception)
}

type authRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return authRepositoryImpl{db: db}
}

func (r authRepositoryImpl) Add(ctx context.Context, newUserData entity.UserWrite) (entity.User, exception.Exception) {
	query := `
		INSERT INTO users(username, password, user_role, created_at, updated_at)
		VALUES
			($1, $2, $3, NOW(), NOW())
		RETURNING id, username, user_role
	`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, newUserData.Username, newUserData.Password, constants.ROLE_USER).Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		return entity.User{}, repository_exception.CreateDBException(err, "auth", "something is wrong in our end")
	}

	return user, nil
}

func (r authRepositoryImpl) FindByUsername(ctx context.Context, username string) (entity.User, exception.Exception) {
	query := `
		SELECT id, username, password, user_role
		FROM users
		WHERE username = $1
		LIMIT 1
	`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		var empty entity.User
		if errors.Is(err, sql.ErrNoRows) {
			return empty, repository_exception.CreateRecordNotFoundException(err, "auth", "record not found")
		}

		return empty, repository_exception.CreateDBException(err, "auth", "something is wrong in our end")
	}

	return user, nil
}

func (r authRepositoryImpl) IsExist(ctx context.Context, username string) (bool, exception.Exception) {
	query := `SELECT 1 FROM users WHERE username = $1`
	res, err := r.db.ExecContext(ctx, query, username)
	if err != nil {
		return false, repository_exception.CreateDBException(err, "auth", "something is wrong in our end")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, repository_exception.CreateDBException(err, "auth", "something is wrong in our end")
	}

	return rowsAffected > 0, nil
}
