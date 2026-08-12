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
)

type CourseItemRepository interface {
	FindByCourse(ctx typing.Context, course entity.Course) ([]entity.CourseItem, exception.Exception)
	AddBulk(ctx typing.Context, items []entity.CourseItem) ([]entity.CourseItem, exception.Exception)
	UpdateSingular(ctx typing.Context, item entity.CourseItem) (entity.CourseItem, exception.Exception)
	DeleteSingular(ctx typing.Context, item entity.CourseItem) (bool, exception.Exception)
}

type courseItemRepositoryImpl struct {
	executor queryexec.QueryExecutor
}

func NewCourseItemRepository(executor queryexec.QueryExecutor) CourseItemRepository {
	return courseItemRepositoryImpl{
		executor: executor,
	}
}

func (r courseItemRepositoryImpl) FindByCourse(ctx typing.Context, course entity.Course) ([]entity.CourseItem, exception.Exception) {
	querystr := `
		SELECT id, name, item_type, item_url, start_time, due_time, description
		FROM course_items
		WHERE
			course_id = $1 AND
			deleted_at IS NULL
	`

	rows, err := r.executor.QueryContext(ctx, querystr, course.ID)
	if err != nil {
		return nil, err
	}

	res, err := querying.ScanForRowsThenTransform(
		rows, dao.NewCourseItemDb,
		func(rows queryexec.Rows, elem dao.CourseItemDb) (dao.CourseItemDb, exception.Exception) {
			scanErr := rows.Scan(
				&elem.ID, &elem.Name, &elem.ItemType, &elem.ItemUrl,
				&elem.StartTime, &elem.DueTime, &elem.Description,
			)

			return elem, scanErr
		},
		converter.CourseItemDbToEntity,
	)

	return res, err
}

func (r courseItemRepositoryImpl) AddBulk(ctx typing.Context, items []entity.CourseItem) ([]entity.CourseItem, exception.Exception) {
	querystr := `
		INSERT INTO course_items(course_id, name, item_type, item_url, start_time, due_time, description, created_at, updated_at)
		VALUES
			%s
		RETURNING id, name, item_type, item_url, start_time, due_time, description
	`

	itemsDb := lambda.MapList(items, converter.CourseItemEntityToDaoDb)

	rowstr, args, _ := querying.BulkValuesFromBeginning(
		itemsDb,
		func(elem dao.CourseItemDb, currentIdx int) querying.DataStrArgs {
			return querying.DataStrArgs{
				RowStr: fmt.Sprintf(
					"(CAST($%d AS BIGINT), $%d, CAST($%d AS SMALLINT), $%d, CAST($%d AS TIMESTAMP), CAST($%d AS TIMESTAMP), $%d, NOW(), NOW())",
					currentIdx, currentIdx+1, currentIdx+2, currentIdx+3, currentIdx+4, currentIdx+5, currentIdx+6,
				),
				Args: []any{elem.Course.ID, elem.Name, elem.ItemType, elem.ItemUrl, elem.StartTime, elem.DueTime, elem.Description},
			}
		},
	)

	querystr = fmt.Sprintf(querystr, rowstr)
	rows, err := r.executor.QueryContext(ctx, querystr, args...)
	if err != nil {
		return nil, err
	}

	res, err := querying.ScanForRowsThenTransform(
		rows, dao.NewCourseItemDb,
		func(rows queryexec.Rows, elem dao.CourseItemDb) (dao.CourseItemDb, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.Name, &elem.ItemType, &elem.ItemUrl, &elem.StartTime, &elem.DueTime, &elem.Description)
			return elem, err
		},
		converter.CourseItemDbToEntity,
	)

	return res, err
}

func (r courseItemRepositoryImpl) UpdateSingular(ctx typing.Context, item entity.CourseItem) (entity.CourseItem, exception.Exception) {
	querystr := `
		UPDATE course_items
		SET
			name = $3,
			item_type = $4,
			item_url = $5,
			start_time = $6,
			due_time = $7,
			description = $8,
			updated_at = NOW()
		WHERE
			id = $1 AND
			course_id = $2 AND
			deleted_at IS NULL
		RETURNING
			id, name, item_type, item_url, start_time, due_time, description
	`

	updateDataDto := converter.CourseItemEntityToDaoDb(item)
	resDto := dao.CourseItemDb{}
	err := r.executor.QueryRowContext(
		ctx, querystr, updateDataDto.ID, updateDataDto.Course.ID,
		updateDataDto.Name, updateDataDto.ItemType, updateDataDto.ItemUrl,
		updateDataDto.StartTime, updateDataDto.DueTime, updateDataDto.Description,
	).Scan(&resDto.ID, &resDto.Name, &resDto.ItemType, &resDto.ItemUrl, &resDto.StartTime, &resDto.DueTime, &resDto.Description)

	if err != nil {
		return entity.CourseItem{}, repository_exception.WrapQueryexecException(err, "course_item")
	}

	res := converter.CourseItemDbToEntity(resDto)
	return res, nil
}

func (r courseItemRepositoryImpl) DeleteSingular(ctx typing.Context, item entity.CourseItem) (bool, exception.Exception) {
	querystr := `
		UPDATE course_items
		SET
			deleted_at = NOW()
		WHERE
			id = $1 AND
			course_id = $2 AND
			deleted_at IS NULL
	`

	res, err := r.executor.ExecContext(ctx, querystr, item.ID, item.Course.ID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
