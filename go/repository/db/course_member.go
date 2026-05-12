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

type CourseMemberRepository interface {
	FindByCourse(ctx typing.Context, course entity.Course) ([]entity.CourseMember, exception.Exception)
	AddMultiple(ctx typing.Context, newMembers []entity.CourseMember) ([]entity.CourseMember, exception.Exception)
	DeleteMultiple(ctx typing.Context, courseId typing.ID, memberIds []typing.ID) (int64, exception.Exception)
	Update(ctx typing.Context, courseId typing.ID, members []entity.CourseMember) ([]entity.CourseMember, exception.Exception)
}

type courseMemberRepositoryImpl struct {
	executor queryexec.QueryExecutor
}

func NewCourseMemberDbRepository(executor queryexec.QueryExecutor) CourseMemberRepository {
	return courseMemberRepositoryImpl{
		executor: executor,
	}
}

func (r courseMemberRepositoryImpl) FindByCourse(ctx typing.Context, course entity.Course) ([]entity.CourseMember, exception.Exception) {
	querystr := `
		SELECT cm.id, u.id, u.email, u.username, cm.course_role
		FROM (
			SELECT id, user_id, course_role
			FROM course_members
			WHERE course_id = $1 AND deleted_at IS NULL
		) AS cm
		JOIN (
			SELECT id, email, username
			FROM users
			WHERE deleted_at IS NULL
		) AS u
		ON cm.user_id = u.id 
	`

	rows, err := r.executor.QueryContext(ctx, querystr, course.ID)
	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "course_member")
	}
	defer r.executor.CloseRows(rows, "course_member - FindByCourse")

	res, err := querying.ScanForRows(
		rows, entity.NewCourseMember,
		func(rows queryexec.Rows, elem entity.CourseMember) (entity.CourseMember, exception.Exception) {
			err = rows.Scan(&elem.ID, &elem.User.ID, &elem.User.Email, &elem.User.Username, &elem.Role)
			return elem, err
		},
	)

	return res, nil
}

func (r courseMemberRepositoryImpl) AddMultiple(ctx typing.Context, newMembers []entity.CourseMember) ([]entity.CourseMember, exception.Exception) {
	querystr := `
		INSERT INTO course_members(course_id, user_id, course_role, created_at, updated_at)
		SELECT *
		FROM (
			VALUES
				%s
		) AS newdata(course_id, user_id, course_role, created_at, updated_at)
		WHERE NOT EXISTS (
			SELECT 1
			FROM course_members
			WHERE course_id = newdata.course_id AND user_id = newdata.user_id AND deleted_at IS NULL
		)
		RETURNING id, user_id, course_role
	`

	paramIdxStart := 1
	rowstr, args, _ := querying.BulkValues(
		newMembers, paramIdxStart,
		func(member entity.CourseMember, currentIdx int) querying.DataStrArgs {
			return querying.DataStrArgs{
				RowStr: fmt.Sprintf(
					"(CAST($%d AS BIGINT), CAST($%d AS BIGINT), CAST($%d AS SMALLINT), NOW(), NOW())",
					currentIdx, currentIdx+1, currentIdx+2,
				),
				Args: []any{member.Course.ID, member.User.ID, member.Role},
			}
		},
	)
	querystr = fmt.Sprintf(querystr, rowstr)
	rows, err := r.executor.QueryContext(ctx, querystr, args...)
	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "course_member")
	}
	defer r.executor.CloseRows(rows, "course member - FindByCourse")

	res, err := querying.ScanForRows(
		rows, entity.NewCourseMember,
		func(rows queryexec.Rows, elem entity.CourseMember) (entity.CourseMember, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.User.ID, &elem.Role)
			return elem, err
		},
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "course_member")
	}

	return res, nil
}

func (r courseMemberRepositoryImpl) DeleteMultiple(ctx typing.Context, courseId typing.ID, memberIds []typing.ID) (int64, exception.Exception) {
	querystr := `
		UPDATE course_members cm
		SET deleted_at = NOW()
		FROM (
			VALUES
			%s
		) AS delete_data(id)
		WHERE
			cm.course_id = $1 AND
			delete_data.id = cm.id AND
			cm.deleted_at IS NULL
	`
	idxStart := 2
	rowstr, args, _ := querying.BulkValues(memberIds, idxStart, func(memberId typing.ID, currentIdx int) querying.DataStrArgs {
		return querying.DataStrArgs{
			RowStr: fmt.Sprintf("(CAST($%d AS BIGINT))", currentIdx),
			Args:   []any{memberId},
		}
	})

	argsInit := []any{courseId}
	args = append(argsInit, args...)

	querystr = fmt.Sprintf(querystr, rowstr)
	res, err := r.executor.ExecContext(ctx, querystr, args...)
	if err != nil {
		return -1, repository_exception.WrapQueryexecException(err, "course_member")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, repository_exception.WrapQueryexecException(err, "course_member")
	}

	return rowsAffected, nil
}

func (r courseMemberRepositoryImpl) Update(ctx typing.Context, courseId typing.ID, members []entity.CourseMember) ([]entity.CourseMember, exception.Exception) {
	querystr := `
		UPDATE course_members cm
		SET
			course_role = update_data.course_role,
			updated_at = NOW()
		FROM (
			VALUES
			%s
		) AS update_data(id, course_role)
		WHERE
			cm.course_id = $1 AND
			update_data.id = cm.id AND
			cm.deleted_at IS NULL
		RETURNING
			cm.id, cm.user_id, cm.course_role
	`

	idxStart := 2
	rowstr, args, _ := querying.BulkValues(
		members, idxStart,
		func(member entity.CourseMember, currentIdx int) querying.DataStrArgs {
			return querying.DataStrArgs{
				RowStr: fmt.Sprintf("(CAST($%d AS BIGINT), CAST($%d AS SMALLINT))", currentIdx, currentIdx+1),
				Args:   []any{member.ID, member.Role},
			}
		},
	)

	querystr = fmt.Sprintf(querystr, rowstr)

	argsInit := []any{courseId}
	args = append(argsInit, args...)

	rows, err := r.executor.QueryContext(ctx, querystr, args...)
	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "course_member")
	}
	defer r.executor.CloseRows(rows, "course_member - Update")

	res, err := querying.ScanForRows(
		rows, entity.NewCourseMember,
		func(rows queryexec.Rows, elem entity.CourseMember) (entity.CourseMember, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.User.ID, &elem.Role)
			return elem, err
		},
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "course_member")
	}

	return res, nil
}
