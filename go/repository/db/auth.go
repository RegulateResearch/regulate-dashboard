package repo_db

import (
	"frascati/comp/queryexec"
	"frascati/constants"
	"frascati/exception"
	"frascati/obj/converter"
	"frascati/obj/dao"
	"frascati/obj/entity"
	repository_exception "frascati/repository/exception"
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

type authRepositoryDbImpl struct {
	executor queryexec.QueryExecutor
}

func NewAuthRepositoryDb(executor queryexec.QueryExecutor) AuthRepository {
	return authRepositoryDbImpl{
		executor: executor,
	}
}

func (r authRepositoryDbImpl) Add(ctx typing.Context, newUserData entity.User) (entity.User, exception.Exception) {
	query := `
		INSERT INTO users(email, username, display_name, password, user_role, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, email, username, user_role
	`

	var user entity.User
	err := r.executor.QueryRowContext(
		ctx, query,
		newUserData.Email, newUserData.Username, newUserData.DisplayName, newUserData.Password, newUserData.Role,
	).Scan(&user.ID, &user.Email, &user.Username, &user.Role)
	if err != nil {
		return entity.User{}, repository_exception.WrapQueryexecException(err, "auth")
	}

	return user, nil
}

func (r authRepositoryDbImpl) FindByEmail(ctx typing.Context, email string) (entity.User, exception.Exception) {
	query := `
		SELECT id, email, username, password, user_role
		FROM users
		WHERE email = $1
		LIMIT 1
	`
	var userDao dao.UserDb
	err := r.executor.QueryRowContext(ctx, query, email).Scan(
		&userDao.ID, &userDao.Email, &userDao.Username, &userDao.Password, &userDao.Role,
	)
	if err != nil {
		return entity.User{}, repository_exception.WrapQueryexecException(err, "auth")
	}

	user := converter.UserDbToEntity(userDao)

	return user, nil
}

func (r authRepositoryDbImpl) IsExistByEmailOrUsername(ctx typing.Context, email string, username string) (bool, exception.Exception) {
	query := `SELECT 1 FROM users WHERE email = $1 OR username = $2`
	res, err := r.executor.ExecContext(ctx, query, email, username)
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "auth")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "auth")
	}

	return rowsAffected > 0, nil
}

func (r authRepositoryDbImpl) FindBySsoData(ctx typing.Context, username string, civitasID string) (entity.User, exception.Exception) {
	query := `
		SELECT id, user_role, has_sso_login
		FROM users
		WHERE username = $1 OR civitas_id = $2
		LIMIT 1
	`

	var userDao dao.UserDb
	err := r.executor.QueryRowContext(ctx, query, username, civitasID).Scan(&userDao.ID, &userDao.Role, &userDao.HasSsoLogin)

	if err != nil {
		return entity.User{}, repository_exception.WrapQueryexecException(err, "auth")
	}

	user := converter.UserDbToEntity(userDao)
	return user, nil
}

func (r authRepositoryDbImpl) AddBySsoData(ctx typing.Context, userData entity.User) (entity.User, exception.Exception) {
	query := `
		INSERT INTO users(username, display_name, civitas_id, user_role, has_sso_login, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, username, display_name, civitas_id, user_role, has_sso_login
	`

	userDao := dao.UserDb{}
	err := r.executor.QueryRowContext(
		ctx, query,
		userData.Username, userData.DisplayName, userData.CivitasID, constants.ROLE_USER, true,
	).Scan(
		&userDao.ID, &userDao.Username, &userDao.DisplayName, &userDao.CivitasID, &userDao.Role, &userDao.HasSsoLogin,
	)

	if err != nil {
		return entity.User{}, repository_exception.WrapQueryexecException(err, "auth")
	}

	user := converter.UserDbToEntity(userDao)
	return user, nil
}

func (r authRepositoryDbImpl) UpdateSsoData(ctx typing.Context, userData entity.User) (bool, exception.Exception) {
	query := `
		UPDATE users
		SET
			civitas_id = $2,
			has_sso_login = $3,
			updated_at = NOW()
		WHERE
			id = $1 AND
			deleted_at IS NULL
	`

	res, err := r.executor.ExecContext(ctx, query, userData.ID, userData.CivitasID, userData.HasSsoLogin)
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "auth")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "auth")
	}

	return rowsAffected > 0, nil
}
