package repo_db

import (
	"frascati/comp/queryexec"
	"frascati/exception"
	"frascati/obj/converter"
	"frascati/obj/dao"
	"frascati/obj/entity"
	repository_exception "frascati/repository/exception"
	"frascati/typing"
	"frascati/utils/querying"
)

type CourseRepository interface {
	Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception)
	FindAll(ctx typing.Context) ([]entity.Course, exception.Exception)
	FindAllByEnrollingUserId(ctx typing.Context, user entity.User) ([]entity.Course, exception.Exception)
	FindById(ctx typing.Context, id typing.ID) (entity.Course, exception.Exception)
	UpdateById(ctx typing.Context, id typing.ID, updateData entity.Course) (bool, exception.Exception)
	DeleteById(ctx typing.Context, id typing.ID) (bool, exception.Exception)
	IsExistById(ctx typing.Context, id typing.ID) (bool, exception.Exception)
}

type courseRepositoryImpl struct {
	executor queryexec.QueryExecutor
}

func NewCourseDbRepository(executor queryexec.QueryExecutor) CourseRepository {
	return courseRepositoryImpl{
		executor: executor,
	}
}

func (r courseRepositoryImpl) Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception) {
	queryStr := `
		INSERT INTO courses(name, course_year, semester, course_url, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, name, course_year, semester, course_url
	`

	var resDb dao.CourseDb

	err := r.executor.
		QueryRowContext(ctx, queryStr, course.Name, course.Year, course.Term, course.Url).
		Scan(&resDb.ID, &resDb.Name, &resDb.Year, &resDb.Term, &resDb.Url)

	res := converter.CourseDbToEntity(resDb)

	if err != nil {
		return entity.Course{}, repository_exception.WrapQueryexecException(err, "course")
	}

	return res, nil
}

func (r courseRepositoryImpl) FindAll(ctx typing.Context) ([]entity.Course, exception.Exception) {
	query := `
		SELECT id, name, course_year, semester, course_url
		FROM courses
		WHERE deleted_at IS NULL
	`

	rows, err := r.executor.QueryContext(ctx, query)
	if err != nil {
		return nil, repository_exception.CreateDBException(err, "courses", "something is wrong in our end")
	}
	defer r.executor.CloseRows(rows, "course - FindAll")

	res, err := querying.ScanForRowsThenTransform(
		rows, dao.NewCourseDb,
		func(rows queryexec.Rows, elem dao.CourseDb) (dao.CourseDb, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.Name, &elem.Year, &elem.Term, &elem.Url)
			return elem, err
		},
		converter.CourseDbToEntity,
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "course")
	}

	return res, nil
}

func (r courseRepositoryImpl) FindAllByEnrollingUserId(ctx typing.Context, user entity.User) ([]entity.Course, exception.Exception) {
	query := `
		SELECT id, name, course_year, semester, course_url
		FROM courses
		WHERE 
			deleted_at IS NULL AND
			EXISTS (
				SELECT 1
				FROM course_members
				WHERE
					user_id = $1 AND
					courses.id = course_members.course_id AND
					deleted_at IS NULL
			)
	`

	rows, err := r.executor.QueryContext(ctx, query, user.ID)
	if err != nil {
		return nil, repository_exception.CreateDBException(err, "courses", "something is wrong in our end")
	}
	defer r.executor.CloseRows(rows, "course - FindAll")

	res, err := querying.ScanForRowsThenTransform(
		rows, dao.NewCourseDb,
		func(rows queryexec.Rows, elem dao.CourseDb) (dao.CourseDb, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.Name, &elem.Year, &elem.Term, &elem.Url)
			return elem, err
		},
		converter.CourseDbToEntity,
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "course")
	}

	return res, nil
}

func (r courseRepositoryImpl) FindById(ctx typing.Context, id typing.ID) (entity.Course, exception.Exception) {
	querystr := `
		SELECT id, name, course_year, semester, course_url
		FROM courses
		WHERE id = $1 AND deleted_at IS NULL
	`

	var courseDb dao.CourseDb
	err := r.executor.QueryRowContext(ctx, querystr, id).Scan(
		&courseDb.ID, &courseDb.Name, &courseDb.Year, &courseDb.Term, &courseDb.Url,
	)

	if err != nil {
		return entity.Course{}, repository_exception.WrapQueryexecException(err, "course")
	}

	res := converter.CourseDbToEntity(courseDb)

	return res, nil
}

func (r courseRepositoryImpl) UpdateById(ctx typing.Context, id typing.ID, updateData entity.Course) (bool, exception.Exception) {
	querystr := `
		UPDATE courses
		SET
			name = $2,
			course_year = $3,
			semester = $4,
			course_url = $5,
			updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	res, err := r.executor.ExecContext(ctx, querystr, id, updateData.Name, updateData.Year, updateData.Term, updateData.Url)
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "course")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "course")
	}

	return rowsAffected > 0, nil
}

func (r courseRepositoryImpl) DeleteById(ctx typing.Context, id typing.ID) (bool, exception.Exception) {
	querystr := `
		UPDATE courses
		SET
			deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	res, err := r.executor.ExecContext(ctx, querystr, id)
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "course")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "course")
	}

	return rowsAffected > 0, nil
}

func (r courseRepositoryImpl) IsExistById(ctx typing.Context, id typing.ID) (bool, exception.Exception) {
	querystr := `
		SELECT 1
		FROM courses
		WHERE id = $1 AND deleted_at IS NULL
	`

	res, err := r.executor.ExecContext(ctx, querystr, id)
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "course")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, repository_exception.WrapQueryexecException(err, "course")
	}

	return rows > 0, nil
}
