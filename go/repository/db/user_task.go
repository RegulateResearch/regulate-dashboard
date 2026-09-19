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

type UserTaskRepository interface {
	FindByUser(ctx typing.Context, user entity.User) ([]entity.UserTask, exception.Exception)
	FindByUserAndId(ctx typing.Context, user entity.User, ID typing.ID) (entity.UserTask, exception.Exception)
	AddInferByNotYetAdded(ctx typing.Context, user entity.User) (dataAffected int64, err exception.Exception)
}

type userTaskRepositoryImpl struct {
	executor queryexec.QueryExecutor
}

func NewUserTaskRepository(executor queryexec.QueryExecutor) UserTaskRepository {
	return userTaskRepositoryImpl{
		executor: executor,
	}
}

func (r userTaskRepositoryImpl) FindByUser(ctx typing.Context, user entity.User) ([]entity.UserTask, exception.Exception) {
	querystr := `
		SELECT
			task_data.id, progress, target_start, actual_start, target_done, actual_done,
			item_data.id, name, start_time, due_time
		FROM (
			SELECT id, item_id, progress, target_start, actual_start, target_done, actual_done
			FROM user_tasks
			WHERE user_id = $1 AND deleted_at IS NULL
		) AS task_data
		JOIN (
			SELECT id, name, start_time, due_time
			FROM course_items
			WHERE deleted_at IS NULL
		) AS item_data ON task_data.item_id = item_data.id
	`

	rows, err := r.executor.QueryContext(ctx, querystr, user.ID)
	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user_task")
	}
	defer r.executor.CloseRows(rows, "user_tasks - FindByUser")

	res, err := querying.ScanForRowsThenTransform(
		rows, dao.NewUserTaskDb,
		func(rows queryexec.Rows, elem dao.UserTaskDb) (dao.UserTaskDb, exception.Exception) {
			err := rows.Scan(
				&elem.ID, &elem.Progress, &elem.TargetStart, &elem.TargetDone, &elem.ActualStart, &elem.ActualDone,
				&elem.Item.ID, &elem.Item.Name, &elem.Item.StartTime, &elem.Item.DueTime,
			)

			return elem, err
		},
		converter.UserTaskDbToEntity,
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "user_task")
	}

	return res, nil
}

func (r userTaskRepositoryImpl) FindByUserAndId(ctx typing.Context, user entity.User, ID typing.ID) (entity.UserTask, exception.Exception) {
	querystr := `
		WITH task_data AS (
			SELECT id, item_id, progress, target_start, actual_start, target_done, actual_done
			FROM user_tasks
			WHERE 
				id = $1 AND
				user_id = $2 AND 
				deleted_at IS NULL
		),
		item_data AS (
			SELECT id, course_id, name, start_time, due_time
			FROM course_items
			WHERE deleted_at IS NULL
		),
		course_data AS (
			SELECT id, name
			FROM courses
			WHERE deleted_at IS NULL
		)
		SELECT
			t.id, t.progress, t.target_start, t.actual_start, t.target_done, t.actual_done,
			i.id, i.name, i.start_time, i.due_time,
			c.id, c.name
		FROM task_data AS t
		JOIN item_data AS i ON t.item_id = i.id
		JOIN course_data AS c ON i.course_id = c.id
	`

	task := dao.UserTaskDb{}
	err := r.executor.QueryRowContext(ctx, querystr).Scan(
		&task.ID, &task.Progress, &task.TargetStart, &task.TargetDone, &task.ActualDone,
		&task.Item.ID, &task.Item.Name, &task.Item.StartTime, &task.Item.DueTime,
		&task.Item.Course.ID, &task.Item.Course.Name,
	)

	if err != nil {
		return entity.UserTask{}, repository_exception.WrapQueryexecException(err, "user_task")
	}

	return converter.UserTaskDbToEntity(task), nil
}

func (r userTaskRepositoryImpl) AddInferByNotYetAdded(ctx typing.Context, user entity.User) (dataAffected int64, err exception.Exception) {
	querystr := `
		INSERT INTO user_tasks(user_id, item_id, created_at, updated_at)
		SELECT $1, item_data.id, NOW(), NOW()
		FROM (
			SELECT course_id
			FROM course_members
			WHERE user_id = $1 AND deleted_at IS NULL
		) AS enroll_data
		JOIN (
			SELECT id, course_id
			FROM course_items
			WHERE deleted_at IS NULL
		) AS item_data ON enroll_data.course_id = item_data.course_id
		WHERE NOT EXISTS (
			SELECT 1
			FROM user_tasks
			WHERE
				user_id = $1 AND
				item_id = item_data.id AND
				deleted_at IS NULL
		)
	`

	res, err := r.executor.ExecContext(ctx, querystr, user.ID)
	if err != nil {
		return int64(-1), repository_exception.WrapQueryexecException(err, "user_task")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return int64(-1), repository_exception.WrapQueryexecException(err, "user_task")
	}

	return rowsAffected, nil
}
