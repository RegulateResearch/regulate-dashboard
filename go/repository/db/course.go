package repo_db

import (
	"frascati/comp/queryexec"
	"frascati/exception"
	"frascati/obj/entity"
	repository_exception "frascati/repository/exception"
	"frascati/typing"
)

type CourseRepository interface {
	Add(ctx typing.Context, course entity.Course) (entity.Course, exception.Exception)
	FindAll(ctx typing.Context) ([]entity.Course, exception.Exception)
	FindById(ctx typing.Context, id typing.ID) (entity.Course, exception.Exception)
	UpdateById(ctx typing.Context, id typing.ID, updateData entity.Course) (bool, exception.Exception)
	DeleteById(ctx typing.Context, id typing.ID) (bool, exception.Exception)
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
	var res entity.Course
	queryStr := `
		INSERT INTO courses(name, course_year, semester, created_at, updated_at)
		VALUES
			($1, $2, $3, NOW(), NOW())
		RETURNING id, name, course_year, semester
	`

	err := r.executor.
		QueryRowContext(ctx, queryStr, course.Name, course.Year, course.Term).
		Scan(&res.ID, &res.Name, &res.Year, &res.Term)

	if err != nil {
		return entity.Course{}, repository_exception.WrapQueryexecException(err, "course")
	}

	return res, nil
}

func (r courseRepositoryImpl) FindAll(ctx typing.Context) ([]entity.Course, exception.Exception) {
	res := make([]entity.Course, 0)
	query := `
		SELECT id, name, course_year, semester
		FROM courses
		WHERE deleted_at IS NULL
	`

	rows, err := r.executor.QueryContext(ctx, query)
	if err != nil {
		return nil, repository_exception.CreateDBException(err, "courses", "something is wrong in our end")
	}
	defer r.executor.CloseRows(rows, "course - FindAll")

	for rows.Next() {
		var course entity.Course
		err := rows.Scan(&course.ID, &course.Name, &course.Year, &course.Term)
		if err != nil {
			return nil, repository_exception.WrapQueryexecException(err, "course")
		}

		res = append(res, course)
	}

	return res, nil
}

func (r courseRepositoryImpl) FindById(ctx typing.Context, id typing.ID) (entity.Course, exception.Exception) {
	var res entity.Course
	querystr := `
		SELECT id, name, course_year, semester
		FROM courses
		WHERE id = $1 AND deleted_at IS NULL
	`

	err := r.executor.QueryRowContext(ctx, querystr, id).Scan(
		&res.ID, &res.Name, &res.Year, &res.Term,
	)

	if err != nil {
		return entity.Course{}, repository_exception.WrapQueryexecException(err, "course")
	}

	return res, nil
}

func (r courseRepositoryImpl) UpdateById(ctx typing.Context, id typing.ID, updateData entity.Course) (bool, exception.Exception) {
	querystr := `
		UPDATE courses
		SET
			name = $2,
			course_year = $3,
			semester = $4,
			updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	res, err := r.executor.ExecContext(ctx, querystr, id, updateData.Name, updateData.Year, updateData.Term)
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
